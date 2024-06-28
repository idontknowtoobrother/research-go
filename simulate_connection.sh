while true; do
  for i in $(seq 1 $((1 + RANDOM % 10))); do
    curl http://localhost:8080/connect
  done
  sleep 23
done