import React, {useState} from "react"
export const Login = (props) => { //parent componenet send some thing to children
    const[email , setEmail] = useState('')
    const[pass , setPass] = useState('')

    const handleSubmit = (e) => {
        e.preventDefault(); // if we dont call page will be reloded and data will be lost
        console.log(email);
    }
    return (

        <div className="auth-form-container">
        <form className="login-form" onSubmit = {handleSubmit}>
            <label htmlFor="email"> email </label>
            <input value={email} onChange={(e)=>setEmail(e.target.value)} type = "email" placeholder="youremail@gmail.com" id = "email" name = "email" />
            <label htmlFor="password">password </label>
            <input value = {pass}  onChange={(e)=>setPass(e.target.value)} type = "password" placeholder="*******" id = "password" name = "password" />
            <button type = "submit">Log In</button>
        </form>

        <button className="link-button" onClick={ () => props.onFormSwitch("register")}>dont have account? click here to sign up</button>
        </div>
    )
  }
  
  

  