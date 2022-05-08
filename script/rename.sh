for i in $(ls ../Solutions/*.c)
do
    echo ${i}
    mv ${i} ${i%.c}.c1
done