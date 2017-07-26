package assistant

type Task struct {
	reader PictureReader
	pipes  []*Pipe
}

func NewTask(reader PictureReader) *Task {
	return &Task{
		reader: reader,
	}
}

func (task *Task) Pipe(converter PictureConverter, writer PictureWriter) *Task {
	task.pipes = append(task.pipes, &Pipe{converter: converter, writer: writer})
	return task
}

func (task *Task) Exec() (pis []*PictureInfo, err error) {
	if len(task.pipes) == 0 {
		return
	}

	buffer, info, err := task.reader.Read()
	if err != nil {
		return nil, err
	}

	var (
		buf  []byte
		errs Errors
	)
	for _, pipe := range task.pipes {
		clone := info.Clone()
		if pipe.converter != NullPictureConvert {
			buf, err = pipe.converter.Convert(buffer, clone)
			if err != nil {
				errs.add(err)
				continue
			}
		} else {
			buf = buffer
		}

		if pipe.writer != nil {
			err = pipe.writer.Write(buf, clone)
			if err != nil {
				errs.add(err)
				continue
			}
		}

		pis = append(pis, clone)
	}

	if len(errs) == 0 {
		return
	}

	return pis, errs
}
