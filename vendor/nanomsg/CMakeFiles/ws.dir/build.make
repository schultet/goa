# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.5

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Remove some rules from gmake that .SUFFIXES does not remove.
SUFFIXES =

.SUFFIXES: .hpux_make_needs_suffix_list


# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/bin/cmake

# The command to remove a file.
RM = /usr/bin/cmake -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg

# Include any dependencies generated for this target.
include CMakeFiles/ws.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/ws.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/ws.dir/flags.make

CMakeFiles/ws.dir/tests/ws.c.o: CMakeFiles/ws.dir/flags.make
CMakeFiles/ws.dir/tests/ws.c.o: tests/ws.c
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object CMakeFiles/ws.dir/tests/ws.c.o"
	/usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -o CMakeFiles/ws.dir/tests/ws.c.o   -c /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/tests/ws.c

CMakeFiles/ws.dir/tests/ws.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/ws.dir/tests/ws.c.i"
	/usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/tests/ws.c > CMakeFiles/ws.dir/tests/ws.c.i

CMakeFiles/ws.dir/tests/ws.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/ws.dir/tests/ws.c.s"
	/usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/tests/ws.c -o CMakeFiles/ws.dir/tests/ws.c.s

CMakeFiles/ws.dir/tests/ws.c.o.requires:

.PHONY : CMakeFiles/ws.dir/tests/ws.c.o.requires

CMakeFiles/ws.dir/tests/ws.c.o.provides: CMakeFiles/ws.dir/tests/ws.c.o.requires
	$(MAKE) -f CMakeFiles/ws.dir/build.make CMakeFiles/ws.dir/tests/ws.c.o.provides.build
.PHONY : CMakeFiles/ws.dir/tests/ws.c.o.provides

CMakeFiles/ws.dir/tests/ws.c.o.provides.build: CMakeFiles/ws.dir/tests/ws.c.o


# Object files for target ws
ws_OBJECTS = \
"CMakeFiles/ws.dir/tests/ws.c.o"

# External object files for target ws
ws_EXTERNAL_OBJECTS =

ws: CMakeFiles/ws.dir/tests/ws.c.o
ws: CMakeFiles/ws.dir/build.make
ws: libnanomsg.so.1.0.0
ws: CMakeFiles/ws.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking C executable ws"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/ws.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/ws.dir/build: ws

.PHONY : CMakeFiles/ws.dir/build

CMakeFiles/ws.dir/requires: CMakeFiles/ws.dir/tests/ws.c.o.requires

.PHONY : CMakeFiles/ws.dir/requires

CMakeFiles/ws.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/ws.dir/cmake_clean.cmake
.PHONY : CMakeFiles/ws.dir/clean

CMakeFiles/ws.dir/depend:
	cd /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/CMakeFiles/ws.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/ws.dir/depend

