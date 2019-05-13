# Install script for directory: /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg

# Set the install prefix
if(NOT DEFINED CMAKE_INSTALL_PREFIX)
  set(CMAKE_INSTALL_PREFIX "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/build")
endif()
string(REGEX REPLACE "/$" "" CMAKE_INSTALL_PREFIX "${CMAKE_INSTALL_PREFIX}")

# Set the install configuration name.
if(NOT DEFINED CMAKE_INSTALL_CONFIG_NAME)
  if(BUILD_TYPE)
    string(REGEX REPLACE "^[^A-Za-z0-9_]+" ""
           CMAKE_INSTALL_CONFIG_NAME "${BUILD_TYPE}")
  else()
    set(CMAKE_INSTALL_CONFIG_NAME "")
  endif()
  message(STATUS "Install configuration: \"${CMAKE_INSTALL_CONFIG_NAME}\"")
endif()

# Set the component getting installed.
if(NOT CMAKE_INSTALL_COMPONENT)
  if(COMPONENT)
    message(STATUS "Install component: \"${COMPONENT}\"")
    set(CMAKE_INSTALL_COMPONENT "${COMPONENT}")
  else()
    set(CMAKE_INSTALL_COMPONENT)
  endif()
endif()

# Install shared libraries without execute permission?
if(NOT DEFINED CMAKE_INSTALL_SO_NO_EXE)
  set(CMAKE_INSTALL_SO_NO_EXE "1")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nanocat.1.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man1" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nanocat.1")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_errno.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_errno.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_strerror.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_strerror.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_symbol.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_symbol.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_symbol_info.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_symbol_info.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_allocmsg.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_allocmsg.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_reallocmsg.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_reallocmsg.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_freemsg.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_freemsg.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_socket.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_socket.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_close.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_close.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_get_statistic.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_get_statistic.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_getsockopt.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_getsockopt.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_setsockopt.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_setsockopt.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_bind.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_bind.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_connect.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_connect.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_shutdown.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_shutdown.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_send.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_send.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_recv.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_recv.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_sendmsg.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_sendmsg.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_recvmsg.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_recvmsg.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_device.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_device.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_cmsg.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_cmsg.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_poll.3.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man3" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_poll.3")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nanomsg.7.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man7" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nanomsg.7")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_pair.7.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man7" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_pair.7")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_reqrep.7.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man7" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_reqrep.7")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_pubsub.7.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man7" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_pubsub.7")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_survey.7.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man7" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_survey.7")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_pipeline.7.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man7" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_pipeline.7")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_bus.7.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man7" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_bus.7")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_inproc.7.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man7" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_inproc.7")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_ipc.7.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man7" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_ipc.7")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_tcp.7.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man7" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_tcp.7")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_ws.7.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man7" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_ws.7")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/doc/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_env.7.html")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/share/man/man7" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nn_env.7")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/include/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/src/nn.h")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/include/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/src/inproc.h")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/include/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/src/ipc.h")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/include/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/src/tcp.h")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/include/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/src/ws.h")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/include/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/src/pair.h")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/include/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/src/pubsub.h")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/include/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/src/reqrep.h")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/include/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/src/pipeline.h")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/include/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/src/survey.h")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/include/nanomsg" TYPE FILE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/src/bus.h")
endif()

if(NOT CMAKE_INSTALL_COMPONENT OR "${CMAKE_INSTALL_COMPONENT}" STREQUAL "Unspecified")
  if(EXISTS "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/bin/nanocat" AND
     NOT IS_SYMLINK "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/bin/nanocat")
    file(RPATH_CHECK
         FILE "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/bin/nanocat"
         RPATH "")
  endif()
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/bin" TYPE EXECUTABLE FILES "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/nanocat")
  if(EXISTS "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/bin/nanocat" AND
     NOT IS_SYMLINK "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/bin/nanocat")
    file(RPATH_CHANGE
         FILE "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/bin/nanocat"
         OLD_RPATH "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg:"
         NEW_RPATH "")
    if(CMAKE_INSTALL_DO_STRIP)
      execute_process(COMMAND "/usr/bin/strip" "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/bin/nanocat")
    endif()
  endif()
endif()

if(NOT CMAKE_INSTALL_LOCAL_ONLY)
  # Include the install script for each subdirectory.
  include("/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/src/cmake_install.cmake")

endif()

if(CMAKE_INSTALL_COMPONENT)
  set(CMAKE_INSTALL_MANIFEST "install_manifest_${CMAKE_INSTALL_COMPONENT}.txt")
else()
  set(CMAKE_INSTALL_MANIFEST "install_manifest.txt")
endif()

string(REPLACE ";" "\n" CMAKE_INSTALL_MANIFEST_CONTENT
       "${CMAKE_INSTALL_MANIFEST_FILES}")
file(WRITE "/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/${CMAKE_INSTALL_MANIFEST}"
     "${CMAKE_INSTALL_MANIFEST_CONTENT}")
