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
include CMakeFiles/cmsg.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/cmsg.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/cmsg.dir/flags.make

CMakeFiles/cmsg.dir/tests/cmsg.c.o: CMakeFiles/cmsg.dir/flags.make
CMakeFiles/cmsg.dir/tests/cmsg.c.o: tests/cmsg.c
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object CMakeFiles/cmsg.dir/tests/cmsg.c.o"
	/usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -o CMakeFiles/cmsg.dir/tests/cmsg.c.o   -c /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/tests/cmsg.c

CMakeFiles/cmsg.dir/tests/cmsg.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/cmsg.dir/tests/cmsg.c.i"
	/usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/tests/cmsg.c > CMakeFiles/cmsg.dir/tests/cmsg.c.i

CMakeFiles/cmsg.dir/tests/cmsg.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/cmsg.dir/tests/cmsg.c.s"
	/usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/tests/cmsg.c -o CMakeFiles/cmsg.dir/tests/cmsg.c.s

CMakeFiles/cmsg.dir/tests/cmsg.c.o.requires:

.PHONY : CMakeFiles/cmsg.dir/tests/cmsg.c.o.requires

CMakeFiles/cmsg.dir/tests/cmsg.c.o.provides: CMakeFiles/cmsg.dir/tests/cmsg.c.o.requires
	$(MAKE) -f CMakeFiles/cmsg.dir/build.make CMakeFiles/cmsg.dir/tests/cmsg.c.o.provides.build
.PHONY : CMakeFiles/cmsg.dir/tests/cmsg.c.o.provides

CMakeFiles/cmsg.dir/tests/cmsg.c.o.provides.build: CMakeFiles/cmsg.dir/tests/cmsg.c.o


# Object files for target cmsg
cmsg_OBJECTS = \
"CMakeFiles/cmsg.dir/tests/cmsg.c.o"

# External object files for target cmsg
cmsg_EXTERNAL_OBJECTS =

cmsg: CMakeFiles/cmsg.dir/tests/cmsg.c.o
cmsg: CMakeFiles/cmsg.dir/build.make
cmsg: libnanomsg.so.1.0.0
cmsg: CMakeFiles/cmsg.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking C executable cmsg"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/cmsg.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/cmsg.dir/build: cmsg

.PHONY : CMakeFiles/cmsg.dir/build

CMakeFiles/cmsg.dir/requires: CMakeFiles/cmsg.dir/tests/cmsg.c.o.requires

.PHONY : CMakeFiles/cmsg.dir/requires

CMakeFiles/cmsg.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/cmsg.dir/cmake_clean.cmake
.PHONY : CMakeFiles/cmsg.dir/clean

CMakeFiles/cmsg.dir/depend:
	cd /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/CMakeFiles/cmsg.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/cmsg.dir/depend

