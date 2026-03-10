import os
import time
import sys

def create_zombie():
    pid = os.fork()

    if pid > 0:
        # Este é o processo PAI
        print(f"[PARENT] Created child with PID: {pid}")
        print("[PARENT] I will sleep for 60s and NOT 'wait' for my child.")
        print("[PARENT] Check 'ps aux' in another terminal to see the <defunct> process.")
        time.sleep(60)
    else:
        # Este é o processo FILHO
        print("[CHILD] I am dying now to become a zombie...")
        sys.exit(0)

if __name__ == "__main__":
    create_zombie()