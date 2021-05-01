package code

/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {

	employeesDict := make(map[int]*Employee, len(employees))

	for _, employee := range employees {
		employeesDict[employee.Id] = employee
	}

	var dfs func(int) int
	dfs = func(id int) int {
		res := employeesDict[id].Importance
		for _, id_ := range employeesDict[id].Subordinates {
			res += dfs(id_)
		}
		return res
	}

	return dfs(id)
}

func GetImportance(employees []*Employee, id int) int{
	return getImportance(employees, id)

}
