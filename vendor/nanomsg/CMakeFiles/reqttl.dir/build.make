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
include CMakeFiles/reqttl.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/reqttl.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/reqttl.dir/flags.make

CMakeFiles/reqttl.dir/tests/reqttl.c.o: CMakeFiles/reqttl.dir/flags.make
CMakeFiles/reqttl.dir/tests/reqttl.c.o: tests/reqttl.c
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object CMakeFiles/reqttl.dir/tests/reqttl.c.o"
	/usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -o CMakeFiles/reqttl.dir/tests/reqttl.c.o   -c /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/tests/reqttl.c

CMakeFiles/reqttl.dir/tests/reqttl.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/reqttl.dir/tests/reqttl.c.i"
	/usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/tests/reqttl.c > CMakeFiles/reqttl.dir/tests/reqttl.c.i

CMakeFiles/reqttl.dir/tests/reqttl.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/reqttl.dir/tests/reqttl.c.s"
	/usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/tests/reqttl.c -o CMakeFiles/reqttl.dir/tests/reqttl.c.s

CMakeFiles/reqttl.dir/tests/reqttl.c.o.requires:

.PHONY : CMakeFiles/reqttl.dir/tests/reqttl.c.o.requires

CMakeFiles/reqttl.dir/tests/reqttl.c.o.provides: CMakeFiles/reqttl.dir/tests/reqttl.c.o.requires
	$(MAKE) -f CMakeFiles/reqttl.dir/build.make CMakeFiles/reqttl.dir/tests/reqttl.c.o.provides.build
.PHONY : CMakeFiles/reqttl.dir/tests/reqttl.c.o.provides

CMakeFiles/reqttl.dir/tests/reqttl.c.o.provides.build: CMakeFiles/reqttl.dir/tests/reqttl.c.o


# Object files for target reqttl
reqttl_OBJECTS = \
"CMakeFiles/reqttl.dir/tests/reqttl.c.o"

# External object files for target reqttl
reqttl_EXTERNAL_OBJECTS =

reqttl: CMakeFiles/reqttl.dir/tests/reqttl.c.o
reqttl: CMakeFiles/reqttl.dir/build.make
reqttl: libnanomsg.so.1.0.0
reqttl: CMakeFiles/reqttl.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking C executable reqttl"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/reqttl.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/reqttl.dir/build: reqttl

.PHONY : CMakeFiles/reqttl.dir/build

CMakeFiles/reqttl.dir/requires: CMakeFiles/reqttl.dir/tests/reqttl.c.o.requires

.PHONY : CMakeFiles/reqttl.dir/requires

CMakeFiles/reqttl.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/reqttl.dir/cmake_clean.cmake
.PHONY : CMakeFiles/reqttl.dir/clean

CMakeFiles/reqttl.dir/depend:
	cd /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/CMakeFiles/reqttl.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/reqttl.dir/depend

