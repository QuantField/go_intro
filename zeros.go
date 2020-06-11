package main

import "fmt"
import "math"



func secant(f func(float64) float64 , a float64, b float64) (float64, int) {
   epsilon, MaxIter := 1.0E-6, 50
   fa,  fb   := f(a),f(b)
   iter := 2
   x2, f2 := b, 0.0
   x1, f1 := a, 0.0
   if (fa*fb>0) {
       fmt.Println("No solution")
	   return 0.0, 0
	}else{
       for (math.Abs((x1-x2)/x2)>epsilon && iter<= MaxIter){
           if (x2<a || x2>b || iter==2) {
               x1,f1 = a,fa
               xm := (a+b)/2
               iter+=1
               fxm := f(xm)
               if (fa*fxm<0) {
                   b,fb = xm,fxm
			   }else{
                   a,fa = xm,fxm
			   }
               x2, f2  = xm, fxm
		   }
           iter+=1
           xv, fv := x2, f2
           x2 += -(x2-x1)*f2/(f2-f1)
           f2 = f(x2)
           x1, f1 = xv, fv
           fmt.Printf("x = %4.8f  f(x) = %4.8f \n", x2, f2)
		 }
		return x2, iter 
	}
   
}


func  brent(f  func(float64) float64, a float64, b float64) (float64, int) {
   epsilon, MaxIter := 1.0E-6, 50
   fa,  fb   := f(a),f(b)
   iter := 2
   
   x1,f1,x2,f2 := a,fa,b,fb
   x3 := x2 + 10
   f3 := 0.0
   if (fa*fb>0){
       print("No solution")
	   return 0.0, 0
   } else{
       for  (math.Abs((x2-x3)/x2)>epsilon && iter<=MaxIter) {
           if (x3<a || x3>b || iter==2){
               x1,f1,x2,f2 = a,fa,b,fb
               xm := (a+b)/2
               iter+=1
               fxm := f(xm)
               if fa*fxm<0 {
                   b,fb = xm,fxm
               }else {
                   a,fa = xm,fxm
			   }
               x3 = (a+b)/2.
               iter+=1
               f3 = f(x3)
			}
           iter+=1
           xv, fv := x3, f3
           x3 = f1*f2*x3/((f3-f2)*(f3-f1)) + 
                f2*f3*x1/((f1-f2)*(f1-f3)) + 
                f3*f1*x2/((f2-f1)*(f2-f3))
           f3 = f(x3)
           x1, f1 = x2, f2
           x2, f2 = xv, fv
		   fmt.Printf("x = %4.8f  f(x) = %4.8f \n", x3, f3)           
		}  
		return x3, iter
    }		
   
}



func main(){

  cubic := func(x float64) float64{
              return x*x*x-2  
	       }
  s,d := -1.0, 3.0		   
  
   secant(cubic,s,d)
   fmt.Println()
   brent(cubic,s,d)

}
