*statement* -> **identifier** = *expression*
| **if** ( *expression* ) *statement*
| **if** ( *expression* ) *statement* **else** *statement*
| **while** ( *expression* ) *statement*
| { *statements* }

*statements* -> *statements* *statement*
| []

*expression* -> *literal*
| **identifier**
| *expression* **operator** *expression*
| ( *expression* )

**operator** -> **arithmetic_operator**
| **relational_operator**
| **logical_operator**

**arithmetic_operator** -> **+** | **-** | **\*** | **/** | **%** | **^**

**relational_operator** -> **==** | **!=** | **<** | **<=** | **>** | **>=**

**logical_operator** -> **&&** | **||** | **!**

*literal* -> **integer** | **float** | **string** | **boolean**

**float** -> **integer** . **integer**

**integer** -> **integer** **digit** | **digit**

**digit** -> 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9

**string** -> " **string** **character** " | " **character** " |
