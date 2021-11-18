package hw06pipelineexecution

import "fmt"

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// init stages
	//var out Out
	executeStage := func(in In, done In, stage Stage, i int) Bi {
		fmt.Printf("Stage init: %v\n", i)
		out := make(Bi, 100)
		tmp := make(Bi, 1)
		var kurwa interface{}
		go func() {
			defer close(out)
			//defer close(tmp)
			for s := range in {
				fmt.Printf("Stage:%v Getting value:%v\n", i, s)
				select {
				case <-done:
					fmt.Printf("Stage closed: %v\n", i)
					return
				case tmp <- s:
					fmt.Printf("Stage:%v changing %v\n", i, s)
					stage(tmp)
					fmt.Printf("Stage-TMP_KURWA:%v changing %v\n", i, s)
					kurwa = <-stage(tmp)
					//<-tmp
					fmt.Printf("Stage:%v sending:%s\n", i, kurwa)
					out <- kurwa
				}
			}
		}()
		return out
	}

	for i, stage := range stages {
		in = executeStage(in, done, stage, i)
		fmt.Printf("Stage: %v\n", i)
	}
	return in

}
