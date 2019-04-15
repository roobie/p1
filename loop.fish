#!/usr/bin/env fish

source functions.fish

function rs
    p1_restart
end

function prompt
    set_color green
    echo -n 'Command> '
    set_color normal
end

echo "Welcome to the development REPL for p1."
echo "N.B. If you exit with ^D, the jobs will be disownded, \
which means you'll have to `kill` them yourself."

while true
    read -p prompt input
    echo $input
    switch $input
    case "rs"
        rs
    case "killall"
        p1_killall
    case "jobs"
        jobs
    case "*"
        echo "Unknown command: '$input'"
    end
end
