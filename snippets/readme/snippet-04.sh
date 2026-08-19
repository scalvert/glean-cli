glean auth login                                  # interactive login (detects the best method automatically)
glean auth login --headless --callback-port 49152 # remote login over an SSH port forward
glean auth status                                 # verify credentials, host, and token expiry
glean auth logout                                 # remove all stored credentials
