package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	rawData := Parse("books_db.flatfile")

	// Run algorithms of our choice against the data
	// bpDChart := amountOfBooksperDecade(rawData)
	// bpAlist := booksPerAuthor(rawData)

	// m is a callable function wrapped in reflect.Value which will accept exactly
	// the parameters as we defined them in the method definition. Note that upon
	// the call of this function, the instance of *Library we passed to reflect.ValueOf()
	// will be used as the receiver type. This means it is important that you already passed
	// the correct instance to the reflect.ValueOf() function.
	library := &Library{}
	m := reflect.ValueOf(library).MethodByName("GetMostSoldBooks")

	// to make m actually callable, we will need to cast it from reflect.Value to an
	// actual function type with the correct signatures. Note how we need to first turn it
	// into an interface type and only then cast it to the function type
	mCallable := m.Interface().(func(GenericParams) Reportable)

	// Making methods generic. the method signatures of the generic mCallable can vary, ie: different parameters
	// We can provide a generic function signature, by crating a wrapper method.

	methodName := "GetMostSoldBooks" // taken from configuration file
	concreteMethod := reflect.ValueOf(library).MethodByName(methodName)

	// Find the wrapper method
	wrapperName := fmt.Sprintf("%sWrap", methodName)
	wrapperMethod := reflect.ValueOf(library).MethodByName(wrapperName)

	// Next we need to access the parameter type passed to the concrete method
	// concreteMethodParamsType will now hold the type of the method parameter struct.
	// For the case of the GetMostSoldBooks this is MostSoldBookParams
	concreteMethodParamsType := concreteMethod.Type().In(0).Elem()

	// In order to be able to retrieve the struct fields (which represent the parameters needed by
	// the analytical algorithm) by their names (which are given in the configuration file),
	// we need to create an instance of the method parameter struct type. We need both a
	// pointer to the instance as well as the instance itself (as will be seen later)
	concreteMethodParamsPtr := reflect.New(concreteMethodParamsType)
	concreteMethodParams := concreteMethodParamsPtr.Elem()

	// At this stage, you can iterate over the keys of the stats element from the
	// configuration file and map the parameter types one-by-one to the fields in the
	// parameter (that is, retrieving the fields of the method parameter struct according to
	// their names). To retrieve a field of a struct by its name, we can use reflect.FieldByName()
	parameterField := concreteMethodParams.FieldByName(configParam)

	// Once we have the parameter fields retrieved, we can map the value given for this parameter
	// in the configuration file to the actual field.
	if configValueInt, isInt := configValue.(int); isInt {
		parameterField.SetInt(int64(configValueInt))
	}

	// The above is for the case of integer values, but we could do the same for each value type
	// we expect (in a trial and error fashion) for each parameter given in the configuration file.
	// Setting the value on the parameter field here will also directly affect the method parameters
	// struct, so we don't need to explicitly alter concreteMethodParams in order to store the parameter
	// value retrieved from the configuration file.

	// Lastly just as we did with the wrapper method, we will cast the concreteMethodParams struct
	// to a GenericParams type. Note that we need to use the pointer type here.
	wrapperParams := concreteMethodParamsPtr.Interface().(GenericParams)

	// Putting it all together
	wrapperMethod.Call(wrapperParams)
}

// It is possible to find a method of a type by its name
type Library struct {
	books []Book
}

// MostSoldBooksParams can contain multiple fields that represents the integer parameters
// that we had defined in the method signature originally.
type MostSoldBooksParams struct {
	startYear int
	endYear   int
}

type Reportable interface {
	Report() HTMLStatisticReport
}

type HTMLStatisticReport struct {
}

type GenericParams interface {
	// IsValid chekds the validity of the parameters passed to the concrete analytical method.
	IsValid() (bool, string)
}

// The return type SoldStat must implement the Reportable interface so that it can be returned
// from the Wrapper method GetMostSoldBooksWrap
func (l *Library) GetMostSoldBooks(p *MostSoldBooksParams) SoldStat {
}

// GetMostSoldBooksWrap is the wrapper method for the concrete method GetMostSoldBooks.
// The parameter must implement GenericParams interface to be able to be called from the concrete method
func (l *Library) GetMostSoldBooksWrap(p GenericParams) Reportable {
	if isValid, reason := p.IsValid(); !isValid {
		log.Fatalf("\nParams invalid:: %s", reason)
		return nil
	}
	return l.GetMostSoldBooks(p.(*MostSoldBooksParams))
}

// implement the GenericParams interface
func (p *MostSoldBooksParams) IsValid() (bool, string) {
	return true, ""
}

func (p SoldStat) Report() HTMLStatisticReport {
	// create report

}
