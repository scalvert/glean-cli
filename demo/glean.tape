Output demo/readme.gif

Set Theme "GitHub Dark"
Set FontFamily "Monaco"
Set TypingSpeed 50ms
Set FontSize 20
Set Width 1200
Set Height 800
Set Padding 20
Set Margin 60
Set MarginFill "#4339F2"
Set Shell "bash"
Set BorderRadius 10

Env "GLEAN_CLI_ANONYMIZE" "1"

# Initial pause
Sleep 500ms

# Show chat functionality
Type "glean chat 'Tell me what Glean Assistant can do. Keep it to 3-4 bullets'" Sleep 500ms Enter
Sleep 15s

Ctrl+L

# Show search functionality
Type "glean search -i 'getting started with glean'" Sleep 500ms Enter
Sleep 10s

# Navigate through results
Type "j" Sleep 500ms
Type "j" Sleep 500ms
Type "j" Sleep 500ms
Type "j" Sleep 500ms
Type "k" Sleep 500ms

# Start a new search
Type "/" Sleep 200ms
Type "glean features" Sleep 500ms Enter
Sleep 5s

# Exit search
Type "q"

Sleep 5s
