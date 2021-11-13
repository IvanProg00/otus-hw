package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	pipeIn := in

	for _, stage := range stages {
		pipeIn = executeStage(pipeIn, done, stage)
	}

	return pipeIn
}

func executeStage(in In, done In, stage Stage) Out {
	out := make(Bi)

	go func() {
		ch := stage(in)
		defer close(out)

		for {
			select {
			case v, ok := <-ch:
				if !ok {
					return
				}
				select {
				case out <- v:
				case <-done:
					return
				}
			case <-done:
				return
			}
		}
	}()

	return out
}
