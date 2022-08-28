package main

import "fmt"

type Student struct {
	id    int
	name  string
	age   int
	class string
}

func NewStudent(id int, name string, age int, class string) *Student {
	return &Student{
		id:    id,
		name:  name,
		age:   age,
		class: class,
	}
}

type StudentClass struct {
	studentList []*Student
}

func NewStudentClass() *StudentClass {
	return &StudentClass{
		studentList: make([]*Student, 0, 100),
	}
}

func (s *StudentClass) addStudent(stu *Student) {
	s.studentList = append(s.studentList, stu)
	fmt.Println("添加学生成功")
}

func (s *StudentClass) stuIdExists(id int) bool {
	flag := false
	for _, value := range s.studentList {
		if value.id == id {
			flag = true
		}
	}
	return flag
}

func (s *StudentClass) updateStudent(stu *Student) {
	stuIdx := -1
	for idx, value := range s.studentList {
		if value.id == stu.id {
			stuIdx = idx
		}
	}
	if stuIdx != -1 {
		s.studentList[stuIdx] = stu
		fmt.Println("更新学生成功")
	} else {
		fmt.Println("更新学生失败，: 学生不存在")
	}
}

func (s *StudentClass) deleteStudent(id int) {
	stuIdx := -1
	for idx, value := range s.studentList {
		if value.id == id {
			stuIdx = idx
		}
	}
	if stuIdx != -1 {
		s.studentList = append(s.studentList[:stuIdx], s.studentList[stuIdx+1])
		fmt.Println("删除学生成功")
	} else {
		fmt.Println("删除学生失败: id 不存在")
	}
}

func (s *StudentClass) displayStudent() {
	for _, student := range s.studentList {
		fmt.Printf("id: %v, name: %s, age: %v, class: %s\n", student.id, student.name, student.age, student.class)
	}
}

func main() {
	studentClass := NewStudentClass()
	student1 := NewStudent(1, "Li Lei", 20, "Class 1")
	student2 := NewStudent(2, "Han Meimei", 20, "Class 1")
	student3 := NewStudent(3, "Tom", 25, "Class 2")

	// Add
	fmt.Println("########## Add ##########")
	studentClass.addStudent(student1)
	studentClass.addStudent(student2)
	studentClass.addStudent(student3)
	studentClass.displayStudent()

	// Update
	fmt.Println("########## Update ##########")
	student1Update := NewStudent(1, "Li Lei", 22, "Class 2")
	studentClass.updateStudent(student1Update)
	studentClass.displayStudent()

	// Delete
	fmt.Println("########## Delete ##########")
	studentClass.deleteStudent(1)
	studentClass.displayStudent()
}
