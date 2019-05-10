export LD_LIBRARY_PATH="$LD_LIBRARY_PATH":vendor/nanomsg/build/lib/
export PYTHONPATH="$PYTHONPATH":vendor/nanomsg4py/
python3 driver/translate.py "${@:1}"
