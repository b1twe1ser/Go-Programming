# BMI Calculator 🧮

#### You have probably heard all about the famous BMI ✨
Just in case you haven't: it's essentially a measure to asses how healthy the weight of a person is.

🚨 *_Disclaimer_* 🚨 The BMI is **not** medically accurate   
and is **not** to be used as measure for ones health 👨🏻‍⚕️   

Why ? Because the weight to height ratio of a person depends on many factors,   
non of which are considered properly. For instance I could be 1.65m tall    
hit the gym all day and weigh 80kg of nearly pure muscle, that doesn't make me overweight     
but the BMI would state that I am overweight 😔

### With that out the way 🚪 let's get started 🎬

#### How the BMI calculator works 💪🏼
BMI = weight(kg) / (height * height) (m^2)



```mermaid
  graph TD;
      Start_Application-->Get_User_Weight;
      Get_User_Weight-->Get_User_Height;
      Get_User_Height-->Apply_Formular;
      Apply_Formular-->Return_Result;
      Return_Result--Optional-->Start_Application;
      
```

#### Range of BMI 📈


| Underweight  | Normal | Overweight | Grade 1 Obesity | Grade 2 Obesity | Grade 3 Obesity |
| ----------- | ------ | ---------- | --------------- | --------------- | --------------- |
| \[0, 18.5)  | \[18.5, 24.9) | \[25, 29.9) | \[30, 34.9) | \[35, 39.9) | \[40, ∞) |

💡 With this list, you can check what category the BMI of the user falls in and return the respective result as well as the BMI, like so 👇🏼   
```go
fmt.Printf("Your BMI is %.2f, your weight is %s \n", bmi, status)
```
