import React, { useState,useRef } from 'react';
import { GoogleMap, useLoadScript, Marker } from '@react-google-maps/api';
import { toast } from "react-toastify";
import DefaultImage from "./a.jpg";
const libraries = ['places'];
const mapContainerStyle = {
  margin : "5%",
  width: '90%',
  height: '50vh',
};
const center = {
  lat: 35.7219, // default latitude
  lng: 51.3347, // default longitude
};


export const Makepost = (props) => {
  const ownerid = props.data;
  const[category , setcategory] = useState('')
  const[title , settitle] = useState('')
  const[area , setarea] = useState('')
  const[price , setprice] = useState('')
  const[numberOfRooms , setnumberOfRooms] = useState('')
  const[yearOfConstruction , setyearOfConstruction] = useState('')
  const[floor , setfloor] = useState('')
  const[description , setdescription] = useState('')
  const[long , setlong] = useState('')
  const[lt , setlt] = useState('')
  const [avatarURL, setAvatarURL] = useState(DefaultImage);
  const [store,setstore]= useState("ندارد")
  const [elevator,setelevator]= useState("ندارد")
  const [parking,setparking]= useState("ندارد")
  const fileUploadRef = useRef();

  const handleImageUpload = (event) => {
    event.preventDefault();
    fileUploadRef.current.click();
  }

  const uploadImageDisplay = () => {
   
      const uploadedFile = fileUploadRef.current.files[0];
      const cachedURL = URL.createObjectURL(uploadedFile);
      setAvatarURL(cachedURL)
  }


    const { isLoaded, loadError } = useLoadScript({
        googleMapsApiKey: 'AIzaSyCNVtpIqQ_AMO-Ov07BInNpgkuinI3Zr6Y',
        libraries,
      });
    
      const [markerPosition, setMarkerPosition] = useState(null);
    
      const handleMapClick = (event) => {
        setlong(event.latLng.lng());
        setlt(event.latLng.lat());
        setMarkerPosition({
          lat: event.latLng.lat(),
          lng: event.latLng.lng(),
          
        });
        
      };
    
      if (loadError) {
        return <div>Error loading maps</div>;
      }
    
      if (!isLoaded) {
        return <div>Loading maps</div>;
      }

      const handleSubmit = (e) => {
        e.preventDefault(); // if we dont call page will be reloded and data will be lost
        let regobj = { category, title ,area  , price,numberOfRooms ,yearOfConstruction ,floor ,description ,store , parking , elevator, lt , long ,ownerid,avatarURL};
            
            //console.log(regobj);
            fetch("http://localhost:8080/api/v1/ads/register", {
                method: "POST",
                headers: { 'content-type': 'application/json' },
                body: JSON.stringify(regobj)
            }).then((res) => {
                toast.success('posted successfully.')
                props.onFormSwitch("mainpage",ownerid)
            }).catch((err) => {
                toast.error('Failed :' + err.message);
            });
        
    }
    
    return (
          <div class="container">
         <div class="toppane">logo will be here</div>
         <div class="d-flex">
         <div class="leftpane">
         <GoogleMap
            mapContainerStyle={mapContainerStyle}
            zoom={10}
            center={center}
            onClick={handleMapClick}
          >
            {markerPosition && <Marker position={markerPosition} />}
          </GoogleMap>
              <div className="relative h-96 w-96 m-8">
      <img 
      style = {{width : "10%"}}
        src={avatarURL}
        alt ="Avatar"
        className="h-96 w-96 rounded-full" />

      <form id="form" encType='multipart/form-data'>
        <button 
          style = {{margin:"1%"}}
          type='submit'
          onClick={handleImageUpload}
        >
          انتخاب تصویر
        </button>
        <input 
          type="file"
          id="file"
          ref={fileUploadRef}
          onChange={uploadImageDisplay}
          hidden />
      </form>  
    </div> 
              <br/>
              <button onClick = {handleSubmit}   style={{marginRight : '15px' , padding : '8px' , backgroundColor : 'lightskyblue'}} >ثبت اگهی</button>
              <button onClick = {() => props.onFormSwitch("mainpage",ownerid)} style={{marginLeft : '15px' , padding : '8px' , backgroundColor : 'lightpink'}}>بازگشت</button>
              
              
              
        </div>
        <div class="middlepane">
          <br/>
          
        <input  onChange={(e)=>settitle(e.target.value)} name="myInput" placeholder="عنوان"/>
        <br></br>
        <br></br>

        <input  onChange={(e)=>setarea(e.target.value)} name="myInput" placeholder="متراژ"/>
        <br></br>
        <br></br>
        <input  onChange={(e)=>setprice(e.target.value)} name="myInput" placeholder="قیمت"/>
        <br></br>
        <br></br>
        <input  onChange={(e)=>setnumberOfRooms(e.target.value)} name="myInput" placeholder="تعداد اتاق"/>
        <br></br>
        <br></br>
        <input  onChange={(e)=>setyearOfConstruction(e.target.value)} name="myInput" placeholder="سال ساخت"/>
        <br></br>
        <br></br>
        <input  onChange={(e)=>setfloor(e.target.value)} name="myInput" placeholder="طبقه"/>
        <br></br>
        <br></br>
        <input  onChange={(e)=>setdescription(e.target.value)} name="myInput" placeholder="سایر توضیحات"/>
        <br></br>
        <br></br>
        </div>
        <div class="rightpane">
        <text>انتخاب دسته بندی اگهی</text>
        <hr className="hr"></hr>
        <br></br>
        <br></br>
        <br></br>

            <input onChange={(e)=>setcategory(e.target.value) } type="radio" value="فروش مسکونی"  />فروش مسکونی
            <br></br>
            <br></br>
            <input onClick={(e)=>setcategory(e.target.value) } type="radio" value="اجاره مسکونی"  />اجاره مسکونی
            <br></br>
            <br></br>

             <input onClick={(e)=>setcategory(e.target.value) } type="radio" value="اجاره تجاری و اداری"  />اجاره تجاری و اداری
             <br></br>
            <br></br>
            <input onClick={(e)=>setcategory(e.target.value) } type="radio" value="فروش تجاری و اداری"   />فروش تجاری و اداری
            <br></br>
            <br></br>
            <input onClick={(e)=>setcategory(e.target.value) } type="radio" value="پروژه های ساخت و ساز"    />پروژه های ساخت و ساز
            <br/>
            <hr className="hr"></hr>
            <input onClick={(e)=>setstore("دارد") } type="radio" value="store"  />انباری
            <br></br>
            <br/>
            <input onClick={(e)=>setelevator("دارد") } type="radio" value="elevator"  />اسانسور
            <br></br>
            <br/>
            <input onClick={(e)=>setparking("دارد") } type="radio" value="parking"  />پارکینگ
            <br></br>
          
        </div>
      </div>
    </div>
        
    )
  }
  
  

  