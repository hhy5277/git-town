package steps

// StepList is a list of steps
// with convenience functions for adding and removing steps.
type StepList struct {
	List []Step
}

// Append adds the given step to the end of this StepList.
func (stepList *StepList) Append(step Step) {
	stepList.List = append(stepList.List, step)
}

// AppendList adds all elements of the given StepList to the end of this StepList.
func (stepList *StepList) AppendList(otherList StepList) {
	stepList.List = append(stepList.List, otherList.List...)
}

// IsEmpty returns whether or not this StepList has any elements
func (stepList *StepList) isEmpty() bool {
	return len(stepList.List) == 0
}

// Peek returns the first element of this StepList.
func (stepList *StepList) Peek() (result Step) {
	if stepList.isEmpty() {
		return nil
	}
	return stepList.List[0]
}

// Pop removes and returns the first element of this StepList.
func (stepList *StepList) Pop() (result Step) {
	if stepList.isEmpty() {
		return nil
	}
	result = stepList.List[0]
	stepList.List = stepList.List[1:]
	return
}

// Prepend adds the given step to the beginning of this StepList.
func (stepList *StepList) Prepend(step Step) {
	stepList.List = append([]Step{step}, stepList.List...)
}

// PrependList adds all elements of the given StepList to the start of this StepList.
func (stepList *StepList) PrependList(otherList StepList) {
	stepList.List = append(otherList.List, stepList.List...)
}