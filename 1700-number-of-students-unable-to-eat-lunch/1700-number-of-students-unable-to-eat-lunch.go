func countStudents(students []int, sandwiches []int) int {
    if len(students) == 0 {
        return 0
    }
    cnt := 0
    for len(students) > 0{
        if students[0] == sandwiches[0] {
            students = dequeue(students)
            sandwiches = dequeue(sandwiches)
            cnt = 0
        } else {
            students = enqueue(students, students[0])
            students = dequeue(students)
            cnt++
            if cnt == len(students) {
                break
            }
        }
    }

    return len(students)
}

func enqueue(queue []int, element int) []int {
    queue = append(queue, element) // Simply append to enqueue.
    return queue
}

func dequeue(queue []int) []int {
    if len(queue) == 1 {
        return []int{}

    }
    return queue[1:] // Slice off the element once it is dequeued.
}