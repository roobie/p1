source env.fish

set base_dir (pwd)
set services authentication authorization accounting

function p1_build
    for service in $services
        cd src/$service/
        echo "Building $service"
        go build
        cd $base_dir
    end
end

function p1_run
  for service in $services
      echo "Running $service"
      src/$service/$service &
  end
end

function p1_killall
    kill (jobs -p)
end

function p1_restart
    p1_killall
    p1_build
    p1_run
end
