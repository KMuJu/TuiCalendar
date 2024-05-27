# TUI Calendar

## Structure

Model---+
        V
    Calendar--+
              V
        MonthTable--|
        WeekTable---|
            Day <---+
            Event <-+

## Ideas
SQLite3 database

## Rendering
Only show the next events

-----------------------|---------------------------------
                       |                                 
Monday 27. may         |       NAME                      
                       |                                 
-----------------------|       DATE                      
12:00-13:00            |                                 
Meating 1              |    Description                  
                       |                                 
-----------------------|                                 
                       |                                 
Friday 31. may         |                                 
                       |                                 
-----------------------|                                 
10:00-10:00            |                                 
Meating 2              |                                 
                       |                                 
-----------------------|                                 
14:15-15:00            |                                 
Meating 3              |                                 
                       |                                 
-----------------------|                                 

....
