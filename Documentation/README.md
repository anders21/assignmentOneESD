# Sarah Anderson's Assignment

## Accuracy of importing:
Using unit tests, I have created a test that imports content from a small file. I test that the content imported is what is expected for the first items, then that the expected number of records have been imported.   
  
## Data valiation
I have identified a situation where the integrity of the data is questionable. The JSON file containing student records and marks have a relationship. The marks belong to a student, and is represented by the foreign key `student_Id`. If a mark is imported, that has a `student_Id` of a student that doesn’t exist, it doesn’t make sense to hold onto this data. My application will import this data, and will display a warning on the console. The mark will still be available for queries.  

## Code decisions:
The time taken to import, unmarshal and report generation is a reflection of many code decisions.  
  
I have choosen to sort the data into a go struct that reflects a relational database table. This means the marks and students are stored separate from each other. The advantage to this is, a flexible querying options. Different queries can be run on the data, and the context of the query will not dramatically change the amount of time to produce a report. For example, If the marks were stored within the student, it becomes labour intensive to collect all the marks from each student, before evening beginning to perform a query.   

I have coded two different ways of generating the report, using the same data struct configurations. One option uses an array, and other uses a map. To find out which is the best, I have conducted a benchmark test and profiled the performance of the two algorithms. 

## How to run performance analysis  

1. `go build -ldflags '-X main.PROFILE="YES"'`   using  powershell
2. rename the file to `profiling` if you like, so you can identify it later  
3. run the file called `assignmentOne.exe`  

## Looking at profiling results
1. `go tool pprof mem.pprof`
2. Run commands like
  1. `top`
  2. `tree`

# How to run in production envrioment  
1. go build  
2. run the file called `assignmentOne.exe`  
  





# Task 2:

-- Read time is quick, but process time is slow