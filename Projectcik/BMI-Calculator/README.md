# BMI Calculator 🧮

#### You have probably heard all about the famous BMI ✨
🚨 *_Disclaimer_* 🚨 The BMI is **not** medically accurate   
and is **not** to be used as measure for ones health 👨🏻‍⚕️

### With that out the way 🚪 let's get started 🎬

#### How the BMI calculator works 💪🏼
BMI = weight(kg) / height * height (m^2)



```mermaid
  graph TD;
      Start_Application-->Get_User_Weight;
      Get_User_Weight-->Get_User_Height;
      Get_User_Height-->Apply_Formular;
      Apply_Formular-->Return_Result;
      Return_Result--Optional-->Start_Application;
      
```

