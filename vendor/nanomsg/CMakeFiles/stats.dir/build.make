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
include CMakeFiles/stats.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/stats.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/stats.dir/flags.make

CMakeFiles/stats.dir/tests/stats.c.o: CMakeFiles/stats.dir/flags.make
CMakeFiles/stats.dir/tests/stats.c.o: tests/stats.c
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object CMakeFiles/stats.dir/tests/stats.c.o"
	/usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -o CMakeFiles/stats.dir/tests/stats.c.o   -c /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/tests/stats.c

CMakeFiles/stats.dir/tests/stats.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/stats.dir/tests/stats.c.i"
	/usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/tests/stats.c > CMakeFiles/stats.dir/tests/stats.c.i

CMakeFiles/stats.dir/tests/stats.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/stats.dir/tests/stats.c.s"
	/usr/bin/cc  $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/tests/stats.c -o CMakeFiles/stats.dir/tests/stats.c.s

CMakeFiles/stats.dir/tests/stats.c.o.requires:

.PHONY : CMakeFiles/stats.dir/tests/stats.c.o.requires

CMakeFiles/stats.dir/tests/stats.c.o.provides: CMakeFiles/stats.dir/tests/stats.c.o.requires
	$(MAKE) -f CMakeFiles/stats.dir/build.make CMakeFiles/stats.dir/tests/stats.c.o.provides.build
.PHONY : CMakeFiles/stats.dir/tests/stats.c.o.provides

CMakeFiles/stats.dir/tests/stats.c.o.provides.build: CMakeFiles/stats.dir/tests/stats.c.o


# Object files for target stats
stats_OBJECTS = \
"CMakeFiles/stats.dir/tests/stats.c.o"

# External object files for target stats
stats_EXTERNAL_OBJECTS =

stats: CMakeFiles/stats.dir/tests/stats.c.o
stats: CMakeFiles/stats.dir/build.make
stats: libnanomsg.so.1.0.0
stats: CMakeFiles/stats.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking C executable stats"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/stats.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/stats.dir/build: stats

.PHONY : CMakeFiles/stats.dir/build

CMakeFiles/stats.dir/requires: CMakeFiles/stats.dir/tests/stats.c.o.requires

.PHONY : CMakeFiles/stats.dir/requires

CMakeFiles/stats.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/stats.dir/cmake_clean.cmake
.PHONY : CMakeFiles/stats.dir/clean

CMakeFiles/stats.dir/depend:
	cd /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg /home/tim/src/gkigit.informatik.uni-freiburg.de/tschulte/dimap/vendor/nanomsg/CMakeFiles/stats.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/stats.dir/depend

