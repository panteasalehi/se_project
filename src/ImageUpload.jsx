import React, {useState, useRef} from 'react'
import DefaultImage from "./download (1).jpg";


const ImageUpload = () => {
  const [avatarURL, setAvatarURL] = useState(DefaultImage);

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

  return (
    <div className="relative h-96 w-96 m-8">
      <img 
      style = {{margin : "2%" ,width : "20%"}}
        src={avatarURL}
        alt ="Avatar"
        className="h-96 w-96 rounded-full" />

      <form id="form" encType='multipart/form-data'>
        <button
          type='submit'
          onClick={handleImageUpload}
          className='flex-center absolute bottom-12 right-14 h-9 w-9 rounded-full'>
          
        </button>
        <input 
          type="file"
          id="file"
          ref={fileUploadRef}
          onChange={uploadImageDisplay}
          hidden />
      </form>  
    </div>
  )
}

export default ImageUpload