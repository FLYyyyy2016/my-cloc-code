# my-cloc-code

这个程序模仿了gocloc，里面的规则从gocloc中拼接而来

通过对文件的遍历，得到了结果，目前只是通过最基本的实现来跑通了功能，后续可能会对准确度和西能进行进一步的修改

``` bash
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