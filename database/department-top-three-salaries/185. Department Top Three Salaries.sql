select d.name as Department, e1.salary as Salary, e1.name as Employee from Employee e1
                                                                               inner join Department d
                                                                                          on d.id = e1.departmentId
where e1.salary in (select top 3 salary from
    (
        select distinct salary, departmentId from Employee
    ) as e2
                    where departmentId = e1.departmentId order by salary desc)
order by salary desc