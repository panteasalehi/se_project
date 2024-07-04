import { Link, useNavigate } from "react-router-dom";
import { toast } from "react-toastify";
import React, {useState} from "react"

export const Register = (props) => {
    const[id , setid] = useState('')
    const[pass , setPass] = useState('')
    const[name , setName] = useState('')
    const[seller , setSeller] = useState('false')
    const[customer , setCustomer] = useState('false')





    
    const handleSubmit = (e) => {
        e.preventDefault(); // if we dont call page will be reloded and data will be lost
<<<<<<< HEAD:frontend/src/Register.jsx
        let regobj = { email, password, name };

        console.log(JSON.stringify(regobj));
        fetch("http://localhost:8080/signup", {
            method: "POST",
            headers: { 'content-type': 'application/json' },
            body: JSON.stringify(regobj),
            credentials: 'include'
        }).then((res) => {
            if (res.ok) {
                toast.success('Registered successfully.');
                props.onFormSwitch("mainpage", email);
            } else {
                return res.json().then(err => {
                    throw new Error(err.message);
                });
            }
        }).catch((err) => {
            toast.error('Failed: ' + err.message);
        });

=======
        let regobj = { id , pass ,name , seller , customer};
            
            //console.log(regobj);
            fetch("http://localhost:8000/user", {
                method: "POST",
                headers: { 'content-type': 'application/json' },
                body: JSON.stringify(regobj)
            }).then((res) => {
                toast.success('Registered successfully.')
                if(customer===true){
                    props.onFormSwitch("mainpage" , id)
                }
                else{
                    props.onFormSwitch("paymentpage" , id)
                }
            }).catch((err) => {
                toast.error('Failed :' + err.message);
            });
        
>>>>>>> 15506f7e1e17e45108bdb91099d4d7f789b1a664:src/Register.jsx
    }
    return (

        <div className="auth-form-container">
        <form className="register-form"  onSubmit = {handleSubmit}>
            <label htmlFor="name"> full name </label>
            <input value={name} onChange={(e)=>setName(e.target.value)} type = "text" placeholder="name" id = "name" name = "name" />
            <label htmlFor="id"> id </label>
            <input value={id} onChange={(e)=>setid(e.target.value) } type = "id" placeholder="youremail@gmail.com" id = "id" name = "id" />
            <label htmlFor="password">password </label>
            <input value = {pass}  onChange={(e)=>setPass(e.target.value)} type = "password" placeholder="*******" id = "password" name = "password" />
            <br></br>
            <div><input onClick={(e)=>setSeller("true") } type="radio" value="parking" />house seller</div>
            <div><input onClick={(e)=>setCustomer("true") } type="radio" value="parking" />customer</div>
            
            

            <button style = {{backgroundColor : "#776854"}} type = "submit">Sign up</button>
            <br/>
        </form>
        <br></br>
        <text onClick={() => props.onFormSwitch("login")}>alreade have account? log in! </text>
        </div>
    )
  }
  
  

  