#!/bin/bash

###########################################################
#
# Run the Go Task List Application as a CLI
# Author : Abhishek Panda
#
###########################################################

API_URL="http://localhost:3000"
GET_ALL_URL=${API_URL}/tasklist/all
GET_ALL_TODO_URL=${API_URL}/tasklist/alltodo
GET_BY_TODAY_URL=${API_URL}/tasklist/today
GET_OVERDUE_URL=${API_URL}/tasklist/overdue
DML_URL=${API_URL}/tasklist
Date=$(date +%Y%m%d)
echo $Date
command=$1

echo -e "[INFO] - `date` - Logs Starting" >> ./shell-logs/tasklist-$Date.logs
logerr=$?
if [[ $logerr -eq 0 ]]
then
    if [ -z "$command" ]
    then
        echo -e "[ERROR] - `date` - command not provided - [FAIL]" >> ./shell-logs/tasklist-$Date.logs
        echo -e "[ERROR] - `date` - command not provided - [FAIL]" 
        exit 10
    fi

    case $command in
    "add")
        read -p "Enter the Task Title : " tasktitle
        read -p "Enter the Due Date for the above Task , Format : {YYYY-MM-DD} : " duedate
        create_task()
        {
            cat <<EOF
              {
                  "TaskTitle":"$tasktitle",
                  "DueDate":"$duedate",
                  "TaskDone":false
              }
EOF
        }
        curl --include -k -X "POST" $DML_URL \
        -H "content-type: application/json" \
        -d "$(create_task)" \
        -o "response.json"

        responseStatusCode=$(cat response.json | grep HTTP/1.1 | cut -d ' ' -f 2)

        if [[ $responseStatusCode -eq 201 ]]
        then
          echo -e "[INFO] - `date` - Task Created Successfully ! - [SUCCESS]"
          echo -e "Response Body : "
          tail -1 response.json | python -m json.tool
          echo -e "[INFO] - `date` - Task Created Successfully ! - [SUCCESS]" >> ./shell-logs/tasklist-$Date.logs
          echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
          tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
	        rm response.json
          exit 51
        else
          echo -e "[Error] - `date` - Error in Task creation !, Status Code : $responseStatusCode - [FAIL]"
          echo -e "Response Body : "
          tail -1 response.json | python -m json.tool
          echo -e "[Error] - `date` - Error in Task creation !, Status Code : $responseStatusCode - [FAIL]" >> ./shell-logs/tasklist-$Date.logs
          echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
          tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
	        rm response.json
          exit 12
        fi
      ;;
    "list")
        read -p "Enter Option number(1. All Tasks , 2. All Todo Tasks , 3. Tasks Due today ,4.Tasks OverDue ,5. Get Task by Title) : " gettype
        case $gettype in
        "1")
            curl --include -k -X "GET" $GET_ALL_URL \
            -H "content-type: application/json" \
            -o "response.json"

            responseStatusCode=$(cat response.json | grep HTTP/1.1 | cut -d ' ' -f 2)

            if [[ $responseStatusCode -eq 200 ]]
            then
              echo -e "[INFO] - `date` - Get Request Successful ! - [SUCCESS]"
              echo -e "Response Body : "
              tail -1 response.json | python -m json.tool
              echo -e "[INFO] - `date` - Get Request Successful ! - [SUCCESS]" >> ./shell-logs/tasklist-$Date.logs
              echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
              tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
              rm response.json
              exit 50
            else
              echo -e "[Error] - `date` - Error in Get Request !, Status Code : $responseStatusCode - [FAIL]"
              echo -e "Response Body : "
              tail -1 response.json | python -m json.tool
              echo -e "[Error] - `date` - Error in Get Request !, Status Code : $responseStatusCode - [FAIL]" >> ./shell-logs/tasklist-$Date.logs
              echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
              tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
              rm response.json
              exit 15
            fi
          ;;
        "2")
            curl --include -k -X "GET" $GET_ALL_TODO_URL \
            -H "content-type: application/json" \
            -o "response.json"

            responseStatusCode=$(cat response.json | grep HTTP/1.1 | cut -d ' ' -f 2)

            if [[ $responseStatusCode -eq 200 ]]
            then
              echo -e "[INFO] - `date` - Get Request Successful ! - [SUCCESS]"
              echo -e "Response Body : "
              tail -1 response.json | python -m json.tool
              echo -e "[INFO] - `date` - Get Request Successful ! - [SUCCESS]" >> ./shell-logs/tasklist-$Date.logs
              echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
              tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
              rm response.json
              exit 50
            else
              echo -e "[Error] - `date` - Error in Get Request !, Status Code : $responseStatusCode - [FAIL]"
              echo -e "Response Body : "
              tail -1 response.json | python -m json.tool
              echo -e "[Error] - `date` - Error in Get Request !, Status Code : $responseStatusCode - [FAIL]" >> ./shell-logs/tasklist-$Date.logs
              echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
              tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
              rm response.json
              exit 15
            fi
          ;;
        "3")
            curl --include -k -X "GET" $GET_BY_TODAY_URL \
            -H "content-type: application/json" \
            -o "response.json"

            responseStatusCode=$(cat response.json | grep HTTP/1.1 | cut -d ' ' -f 2)

            if [[ $responseStatusCode -eq 200 ]]
            then
              echo -e "[INFO] - `date` - Get Request Successful ! - [SUCCESS]"
              echo -e "Response Body : "
              tail -1 response.json | python -m json.tool
              echo -e "[INFO] - `date` - Get Request Successful ! - [SUCCESS]" >> ./shell-logs/tasklist-$Date.logs
              echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
              tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
              rm response.json
              exit 50
            else
              echo -e "[Error] - `date` - Error in Get Request !, Status Code : $responseStatusCode - [FAIL]"
              echo -e "Response Body : "
              tail -1 response.json | python -m json.tool
              echo -e "[Error] - `date` - Error in Get Request !, Status Code : $responseStatusCode - [FAIL]" >> ./shell-logs/tasklist-$Date.logs
              echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
              tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
              rm response.json
              exit 15
            fi
          ;;
        "4")
            curl --include -k -X "GET" $GET_OVERDUE_URL \
            -H "content-type: application/json" \
            -o "response.json"

            responseStatusCode=$(cat response.json | grep HTTP/1.1 | cut -d ' ' -f 2)

            if [[ $responseStatusCode -eq 200 ]]
            then
              echo -e "[INFO] - `date` - Get Request Successful ! - [SUCCESS]"
              echo -e "Response Body : "
              tail -1 response.json | python -m json.tool
              echo -e "[INFO] - `date` - Get Request Successful ! - [SUCCESS]" >> ./shell-logs/tasklist-$Date.logs
              echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
              tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
              rm response.json
              exit 50
            else
              echo -e "[Error] - `date` - Error in Get Request !, Status Code : $responseStatusCode - [FAIL]"
              echo -e "Response Body : "
              tail -1 response.json | python -m json.tool
              echo -e "[Error] - `date` - Error in Get Request !, Status Code : $responseStatusCode - [FAIL]" >> ./shell-logs/tasklist-$Date.logs
              echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
              tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
              rm response.json
              exit 15
            fi
          ;;
        "5")
            read -p "Enter Task Title you want to fetch : " querytasktitle
            querytasktitle=$( printf "%s\n" "$querytasktitle" | sed 's/ /%20/g' )
            GET_BY_TITLE_URL=${DML_URL}/${querytasktitle}
            curl --include -k -X "GET" $GET_BY_TITLE_URL \
            -H "content-type: application/json" \
            -o "response.json"

            responseStatusCode=$(cat response.json | grep HTTP/1.1 | cut -d ' ' -f 2)

            if [[ $responseStatusCode -eq 200 ]]
            then
              echo -e "[INFO] - `date` - Get Request Successful ! - [SUCCESS]"
              echo -e "Response Body : "
              tail -1 response.json | python -m json.tool
              echo -e "[INFO] - `date` - Get Request Successful ! - [SUCCESS]" >> ./shell-logs/tasklist-$Date.logs
              echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
              tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
              rm response.json
              exit 50
            else
              echo -e "[Error] - `date` - Error in Get Request !, Status Code : $responseStatusCode - [FAIL]"
              echo -e "Response Body : "
              tail -1 response.json | python -m json.tool
              echo -e "[Error] - `date` - Error in Get Request !, Status Code : $responseStatusCode - [FAIL]" >> ./shell-logs/tasklist-$Date.logs
              echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
              tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
              rm response.json
              exit 15
            fi
          ;;
        *)
            echo -e "[Error] - `date` - Wrong Option chosen in the Task : List Menu !- [FAIL]"
            echo -e "[Error] - `date` - Wrong Option chosen in the Task : List Menu !- [FAIL]" >> ./shell-logs/tasklist-$Date.logs
          ;;
        esac
      ;;
    "done")
        read -p "Enter Task Title you wish to mark complete : " querytasktitle
        tasktitle=${querytasktitle}
        querytasktitle=$( printf "%s\n" "$querytasktitle" | sed 's/ /%20/g' )
        GET_BY_TITLE_URL=${DML_URL}/${querytasktitle}
        curl --include -k -X "GET" $GET_BY_TITLE_URL \
        -H "content-type: application/json" \
        -o "response.json"
        responseStatusCode=$(cat response.json | grep HTTP/1.1 | cut -d ' ' -f 2)
        if [[ $responseStatusCode -eq 200 ]]
        then
          dueDate=$(tail -1 response.json | python -c "import sys,json; print json.load(sys.stdin)['dueDate']")
          rm response.json
          done_task()
          {
              cat <<EOF
                {
                    "TaskTitle":"$tasktitle",
                    "DueDate":"$dueDate",
                    "TaskDone":true
                }
EOF
          }
          cp ./database/tasklist.db ./database/tasklist_temp.db
          rm -f ./database/tasklist.db
          mv ./database/tasklist_temp.db ./database/tasklist.db


          curl --include -k -X "PUT" $DML_URL \
          -H "content-type: application/json" \
          -d "$(done_task)" \
          -o "response.json"

          responseStatusCode=$(cat response.json | grep HTTP/1.1 | cut -d ' ' -f 2)

          if [[ $responseStatusCode -eq 200 ]]
          then
            echo -e "[INFO] - `date` - Task Marked Done ! - [SUCCESS]"
            echo -e "Response Body : "
            tail response.json
            tail -1 response.json
            tail -1 response.json | python -m json.tool
            echo -e "[INFO] - `date` - Task Marked Done ! - [SUCCESS]" >> ./shell-logs/tasklist-$Date.logs
            echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
            tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
            rm response.json
            exit 52
          else
            echo -e "[Error] - `date` - Error in marking Task Done !, Status Code : $responseStatusCode - [FAIL]"
            echo -e "Response Body : "
            tail -1 response.json | python -m json.tool
            echo -e "[Error] - `date` - Error in marking Task Done !, Status Code : $responseStatusCode - [FAIL]" >> ./shell-logs/tasklist-$Date.logs
            echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
            tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
            rm response.json
            exit 14
          fi

        else
          echo -e "[Error] - `date` - Error in Get Request inside Done Call !, Status Code : $responseStatusCode - [FAIL]"
          echo -e "Response Body : "
          tail -1 response.json | python -m json.tool
          echo -e "[Error] - `date` - Error in Get Request inside Done Call !, Status Code : $responseStatusCode - [FAIL]" >> ./shell-logs/tasklist-$Date.logs
          echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
          tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
          rm response.json
          exit 15
        fi
      ;;
    "updateduedate")
        read -p "Enter Task Title you wish to update the duedate : " querytasktitle
        read -p "Enter the new Due Date : " duedate
        tasktitle=${querytasktitle}
        querytasktitle=$( printf "%s\n" "$querytasktitle" | sed 's/ /%20/g' )
        GET_BY_TITLE_URL=${DML_URL}/${querytasktitle}
        curl --include -k -X "GET" $GET_BY_TITLE_URL \
        -H "content-type: application/json" \
        -o "response.json"
        responseStatusCode=$(cat response.json | grep HTTP/1.1 | cut -d ' ' -f 2)
        if [[ $responseStatusCode -eq 200 ]]
        then
          taskDone=$(tail -1 response.json | python -c "import sys,json; print json.load(sys.stdin)['taskDone']")
          if [ "$taskDone" = "True" ]
          then
            taskDone=true
          elif [ "$taskDone" = "False" ]
          then
            taskDone=false
          fi
          rm response.json
          upadate_duedate_task()
          {
              cat <<EOF
                {
                    "TaskTitle":"$tasktitle",
                    "DueDate":"$duedate",
                    "TaskDone":$taskDone
                }
EOF
          }
          cp ./database/tasklist.db ./database/tasklist_temp.db
          rm -f ./database/tasklist.db
          mv ./database/tasklist_temp.db ./database/tasklist.db

          curl --include -k -X "PUT" $DML_URL \
          -H "content-type: application/json" \
          -d "$(upadate_duedate_task)" \
          -o "response.json"

          responseStatusCode=$(cat response.json | grep HTTP/1.1 | cut -d ' ' -f 2)

          if [[ $responseStatusCode -eq 200 ]]
          then
            echo -e "[INFO] - `date` - Task Due Date Updated ! - [SUCCESS]"
            echo -e "Response Body : "
            tail -1 response.json | python -m json.tool
            echo -e "[INFO] - `date` - Task Due Date Updated ! - [SUCCESS]" >> ./shell-logs/tasklist-$Date.logs
            echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
            tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
            rm response.json
            exit 52
          else
            echo -e "[Error] - `date` - Error in Updating Task's Due Date !, Status Code : $responseStatusCode - [FAIL]"
            echo -e "Response Body : "
            tail -1 response.json | python -m json.tool
            echo -e "[Error] - `date` - Error in Updating Task's Due Date !, Status Code : $responseStatusCode - [FAIL]" >> ./shell-logs/tasklist-$Date.logs
            echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
            tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
            rm response.json
            exit 14
          fi

        else
          echo -e "[Error] - `date` - Error in Get Request inside Update Due Date Call !, Status Code : $responseStatusCode - [FAIL]"
          echo -e "Response Body : "
          tail -1 response.json | python -m json.tool
          echo -e "[Error] - `date` - Error in Get Request inside Update Due Date Call !, Status Code : $responseStatusCode - [FAIL]" >> ./shell-logs/tasklist-$Date.logs
          echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
          tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
          rm response.json
          exit 15
        fi
      ;;
    "delete")
        read -p "Enter the Task Title to delete: " tasktitle
        delete_task()
        {
            cat <<EOF
              {
                  "TaskTitle":"$tasktitle",
                  "DueDate":" ",
                  "TaskDone":false
              }
EOF
        }
        curl --include -k -X "DELETE" $DML_URL \
        -H "content-type: application/json" \
        -d "$(delete_task)" \
        -o "response.json"

        responseStatusCode=$(cat response.json | grep HTTP/1.1 | cut -d ' ' -f 2)

        if [[ $responseStatusCode -eq 200 ]]
        then
          echo -e "[INFO] - `date` - Task Deleted Successfully ! - [SUCCESS]"
          echo -e "Response Body : "
          tail -1 response.json | python -m json.tool
          echo -e "[INFO] - `date` - Task Deleted Successfully ! - [SUCCESS]" >> ./shell-logs/tasklist-$Date.logs
          echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
          tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
	        rm response.json
          exit 53
        else
          echo -e "[Error] - `date` - Error in Task deletion !, Status Code : $responseStatusCode - [FAIL]"
          echo -e "Response Body : "
          tail -1 response.json | python -m json.tool
          echo -e "[Error] - `date` - Error in Task deletion !, Status Code : $responseStatusCode - [FAIL]" >> ./shell-logs/tasklist-$Date.logs
          echo -e "Response Body : " >> ./shell-logs/tasklist-$Date.logs
          tail -1 response.json | python -m json.tool >> ./shell-logs/tasklist-$Date.logs
	        rm response.json
          exit 16
        fi
      ;;
    *)
        echo -e "[Error] - `date` - Wrong Command provided , Please check (add / list / done /updateduedate / delete) !- [FAIL]"
        echo -e "[Error] - `date` - Wrong Command provided , Please check (add / list / done /updateduedate / delete) !- [FAIL]" >> ./shell-logs/tasklist-$Date.logs
      ;;
    esac

else
    echo -e "[ERROR] - `date` - Logs could not be created . logs folder not found in current directory - [FAIL]"
    exit 11
fi