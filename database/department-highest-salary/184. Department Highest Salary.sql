SELECT dname as Department, ename as Employee, esalary as Salary FROM (
  SELECT d.name as dname, e.name as ename, e.salary as esalary, depId, maxSal.sal, e.departmentId as edepId
  FROM Employee as e INNER JOIN Department as d ON (e.departmentId = d.id),
       (
           SELECT MAX(salary) as sal, e.departmentId as depId FROM Employee as e GROUP BY departmentId
       ) as maxSal
  HAVING Salary = maxSal.sal AND e.departmentId = depId
) as resultTable
GROUP BY Employee