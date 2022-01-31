#/bin/bash

for i in {1..5}
do
   curl http://localhost:8080/public/health-check
   if [ $? == 0 ]
   then
      echo "Connection ready"
      exit 0
   fi
   echo "Check #${i} failed"
   echo "Retry in 3s..."
   sleep 3
done 

echo "All checks failed."
exit 1 