/*
*	4) Реализовать постоянную запись данных в канал (главный поток).
*	Реализовать набор из N воркеров,
*	которые читают произвольные данные из канала и выводят в stdout.
*	Необходима возможность выбора количества воркеров при старте.
*
*	Программа должна завершаться по нажатию Ctrl+C.
*	Выбрать и обосновать способ завершения работы всех воркеров.
*
*********************************************************************/
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

// worker выводит в stdout данные из канала
func worker(ch <-chan string, id int) {
	// цикл прекратится при закрытии канала
	for msg := range ch {
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%s Done [worker-%d]\n", msg, id)
	}
}

func main() {
	var grs int // количество воркеров

	numCpu := runtime.NumCPU()
	maxGrs := numCpu*256 - 1 // максимальное количество воркеров

	// возможность выбора количества воркеров при старте
	if len(os.Args) > 1 {
		// нас интересует только первый аргумент,
		// возможные остольные просто игнорируем
		arg, err := strconv.Atoi(os.Args[1])

		// если аргумент не конвертируется в int
		// или меньше одного
		// или больше maxGrs
		if err != nil || arg < 1 || arg > maxGrs {
			log.Fatalf("argument: %s expected from 1 to %d", os.Args[1], maxGrs)
		}

		// устанавливаем количество воркеров
		grs = arg
		fmt.Printf("\nrunning workers %d \n", grs)

	} else {
		// если аргумента не было
		// устанавливаем количество воркеров равное NumCPU
		grs = numCpu
		fmt.Printf("\nrunning workers %d (default NumCPU)\n", grs)

		// сообщаем о возможности запуска с указанием
		// количества воркеров
		fmt.Printf("\nyou can specify the quantity workers from 1 to %d\n\nexemple go run path 2\n\n", maxGrs)
	}

	ch := make(chan string)

	for i := 0; i < grs; i++ {
		go worker(ch, i+1 /* worker id */)
	}

	signals := make(chan os.Signal, 2)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)

	fmt.Print("press Ctrl+C to stop \n\n")
	time.Sleep(1000 * time.Millisecond)

	taskId := 1

	for {
		select {
		case <-signals:
			// завершения работы по нажатию Ctrl+C 
			//
			// даем сигнал воркерам
			close(ch)
			fmt.Print(" stoping\n")

			// с одной стороны мы могли бы подаждать пока
			// все уже отправленные задачи воркерам завершатся
			// 
			// с другой стороны пользователь ожидает завершение 
			// немедленно
            
			// зависит от конкретного случая использования

			// идем на компромис
			// ждем разумное время, кто успел тот успел
			time.Sleep(200 * time.Millisecond)
			return

		default:
			fmt.Printf("Add [task-%d]\n", taskId)
			ch <- "Task"
			taskId += 1
		}
	}
}
