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
include CMakeFiles/list.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/list.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/list.dir/flags.make

CMakeFiles/list.dir/tests/list.c.o: CMakeFiles/list.dir/flags.make
CMakeFiles/list.dir/tests/list.c.o: tests/list.c
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object CMakeFiles/list.dir/tests/list.c.o"
	/usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -o CMakeFiles/list.dir/tests/list.c.o   -c /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/tests/list.c

CMakeFiles/list.dir/tests/list.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/list.dir/tests/list.c.i"
	/usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/tests/list.c > CMakeFiles/list.dir/tests/list.c.i

CMakeFiles/list.dir/tests/list.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/list.dir/tests/list.c.s"
	/usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/tests/list.c -o CMakeFiles/list.dir/tests/list.c.s

CMakeFiles/list.dir/tests/list.c.o.requires:

.PHONY : CMakeFiles/list.dir/tests/list.c.o.requires

CMakeFiles/list.dir/tests/list.c.o.provides: CMakeFiles/list.dir/tests/list.c.o.requires
	$(MAKE) -f CMakeFiles/list.dir/build.make CMakeFiles/list.dir/tests/list.c.o.provides.build
.PHONY : CMakeFiles/list.dir/tests/list.c.o.provides

CMakeFiles/list.dir/tests/list.c.o.provides.build: CMakeFiles/list.dir/tests/list.c.o


# Object files for target list
list_OBJECTS = \
"CMakeFiles/list.dir/tests/list.c.o"

# External object files for target list
list_EXTERNAL_OBJECTS =

list: CMakeFiles/list.dir/tests/list.c.o
list: CMakeFiles/list.dir/build.make
list: libnanomsg.so.1.0.0
list: CMakeFiles/list.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking C executable list"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/list.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/list.dir/build: list

.PHONY : CMakeFiles/list.dir/build

CMakeFiles/list.dir/requires: CMakeFiles/list.dir/tests/list.c.o.requires

.PHONY : CMakeFiles/list.dir/requires

CMakeFiles/list.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/list.dir/cmake_clean.cmake
.PHONY : CMakeFiles/list.dir/clean

CMakeFiles/list.dir/depend:
	cd /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/CMakeFiles/list.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/list.dir/depend

