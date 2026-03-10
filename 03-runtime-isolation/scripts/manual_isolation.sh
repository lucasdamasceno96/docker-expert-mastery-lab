#!/bin/bash

# Configuration
CG_NAME="mastery_limit"
CG_PATH="/sys/fs/cgroup/$CG_NAME"

echo "[1/3] Setting up Cgroup v2 memory limit..."
# Ensure we are root (inside the privileged container)
if [ "$EUID" -ne 0 ]; then 
  echo "Please run as root"
  exit
fi

mkdir -p "$CG_PATH"
# 50MB limit in bytes
echo "52428800" > "$CG_PATH/memory.max" 

echo "[2/3] Spawning isolated shell (UTS + PID namespaces)..."
# The --fork is vital to make the shell the PID 1 of its own world
# The --mount is added to ensure our /proc remount doesn't leak to the host
unshare --uts --pid --mount --fork /bin/bash -c "
    hostname 'isolated-shell'
    
    # Attach the child shell to the cgroup
    echo \$$ > $CG_PATH/cgroup.procs
    
    # The Senior Touch: mount a private /proc so 'ps' works
    # We use a fresh mount to isolate process visibility
    mount -t proc proc /proc
    
    echo '[OK] You are now in a manual container!'
    echo \"Your PID according to this namespace: \$$\"
    echo \"Your Hostname: \$(hostname)\"
    echo \"Memory Limit (Cgroup): \$(cat $CG_PATH/memory.max) bytes\"
    echo '----------------------------------------------------------'
    
    # Start the interactive session
    /bin/bash
    
    # Cleanup when leaving the sub-shell
    echo '[3/3] Cleaning up...'
    umount /proc
"

echo "[Finished] You have exited the isolated process."