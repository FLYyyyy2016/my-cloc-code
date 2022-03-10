# my-cloc-code

这个程序模仿了gocloc，里面的规则从gocloc中拼接而来

通过对文件的遍历，得到了结果，目前只是通过最基本的实现来跑通了功能，后续可能会对准确度和西能进行进一步的修改

my-cloc-code:
```bash
time ./my-cloc-code ~/go/go1.12.17/src
INFO[0000] ------------------------------------         
INFO[0000]   language   file    blank   comment code           
INFO[0000]         Go   4085    130708  115345  1519100        
INFO[0000]   Assembly   526     10801   17647   101575            
INFO[0000] Plain Text   268     2666    0       20020                  
INFO[0000]          C   67      720     369     4872                   
INFO[0000]       Perl   14      340     208     2339                   
INFO[0000]       JSON   11      0       0       2042                       
INFO[0000] Bourne Shell 8       137     585     1054                  
INFO[0000]   Markdown   6       227     0       706                       
INFO[0000]       BASH   14      120     232     543                    
INFO[0000]     Python   1       129     52      408                      
INFO[0000]      Batch   5       52      1       241                        
INFO[0000]   C Header   11      56      162     216                     
INFO[0000]        Awk   1       1       6       7                           
INFO[0000]        CSS   1       0       0       1                           
INFO[0000]       HTML   1       0       0       1                           
INFO[0000] ------------------------------------
./my-cloc-code ~/go/go1.12.17/src  0.48s user 0.06s system 80% cpu 0.677 total
```

gocloc:
```bash
time ./bin/gocloc ~/go/go1.12.17/src
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                            4027         133001         229008        1396348
Assembly                       524          11880          17378         100718
Plain Text                     266           2648              0          19943
C                               65            729            590           4627
JSON                            11              0              0           2042
BASH                            20            246            834           1537
Perl                            15            355           1519           1108
Markdown                         6            227              0            706
Python                           1            130            101            358
Batch                            5             52              1            241
C Header                        11             56            162            216
Plan9 Shell                      4             22             48             86
Bourne Shell                     3             12              8             36
Awk                              1              1              6              7
Makefile                         4              3              7              7
CSS                              1              0              0              1
HTML                             1              0              0              1
-------------------------------------------------------------------------------
TOTAL                         4965         149362         249662        1527982
-------------------------------------------------------------------------------
./bin/gocloc ~/go/go1.12.17/src  0.48s user 0.14s system 95% cpu 0.661 total
```
cloc:

```bash
time cloc ~/go/go1.12.17/src 
    5518 text files.
    5407 unique files.                                          
     822 files ignored.

1 error:
Line count, exceeded timeout:  /home/liufei/go/go1.12.17/src/internal/x/net/idna/tables.go

github.com/AlDanial/cloc v 1.82  T=13.11 s (358.4 files/s, 145219.9 lines/s)
-----------------------------------------------------------------------------------
Language                         files          blank        comment           code
-----------------------------------------------------------------------------------
Go                                4027         133001         209621        1415735
Assembly                           524          11869          17378         100729
C                                   65            729            589           4628
Perl                                15            355            369           2258
JSON                                11              0              0           2042
Bourne Shell                         9            138            608           1032
Markdown                             6            227              0            706
Bourne Again Shell                  14            120            234            541
Python                               1            130            101            358
DOS Batch                            5             52              1            241
C/C++ Header                        11             56            162            216
Windows Resource File                4             22              0            134
RobotFramework                       1              0              0            106
awk                                  1              1              6              7
make                                 4              3              7              7
CSS                                  1              0              0              1
HTML                                 1              0              0              1
-----------------------------------------------------------------------------------
SUM:                              4700         146703         229076        1528742
-----------------------------------------------------------------------------------
cloc ~/go/go1.12.17/src  12.70s user 0.22s system 97% cpu 13.286 total
```
