from pynng import Pair1


class AgentComm:

    def __init__(self, agent_id, agent_url):
        self.recv = Pair1(recv_timeout=None, send_timeout=None, polyamorous=True)
        self.recv.listen(agent_url[agent_id])
        self.send = []
        for i, url in enumerate(agent_url):
            if i != agent_id:
                sock = Pair1(recv_timeout=None, send_timeout=None, polyamorous=True)
                sock.dial(url)
                self.send += [sock]
            else:
                self.send += [None]

        self.is_master = (agent_id == 0)
        self.agent_id = agent_id
        self.agent_size = len(agent_url)

        self.msgs = []

    def sendToAll(self, data):
        import marshal as pickle
        buf = pickle.dumps((self.agent_id, data), -1)
        #buf = memoryview(buf)
        for sock in self.send:
            if sock is not None:
                sock.send(buf)

    def _sendTo(self, data, agent_id):
        import marshal as pickle
        buf = pickle.dumps((self.agent_id, data), -1)
        #buf = memoryview(buf)
        #try:
        self.send[agent_id].send(buf)
        #except:
        #    print("Timeout on _sendTo")
        #    pass

    def sendInRing(self, data):
        ai = (self.agent_id + 1) % self.agent_size
        return self._sendTo(data, ai)

    def sendToMaster(self, data):
        return self._sendTo(data, 0)

    def _recv(self):
        import marshal as pickle
        buf = self.recv.recv()
        return pickle.loads(buf)

    def recvFromMaster(self):
        return self.recvFrom(0)

    def recvFrom(self, src_id):
        for i, d in enumerate(self.msgs):
            if d[0] == src_id:
                self.msgs.pop(i)
                return d[1]
        d = self._recv()
        while d[0] != src_id:
            self.msgs.append(d)
            d = self._recv()
        return d[1]

    def recvFromAll(self):
        res = []
        for i in range(self.agent_size):
            if i == self.agent_id:
                continue
            res += [self.recvFrom(i)]
        return res

    def recvInRing(self):
        src_id = self.agent_id - 1
        if src_id < 0:
            src_id = self.agent_size - 1
        return self.recvFrom(src_id)

    def close(self):
        if self.is_master:
            self.sendToAll(None)
            self.recvFromAll()
        else:
            self.recvFromMaster()
            self.sendToMaster(None)

        #self.recv.shutdown()
        self.recv.close()
        for sock in self.send:
            if sock is not None:
                #sock.shutdown()
                sock.close()
