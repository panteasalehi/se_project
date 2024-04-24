import React, {useState} from "react"

export const Mainpage = (probs) => {
    return (
        <view >
            <view >
                <button className = "newpost-button">پروفایل من</button>
            </view>
            <view >
                <button className = "myprofile-button" onClick={() => probs.onFormSwitch("makepost")}>ثبت اگهی</button>
            </view>
        </view>
    )
  }
  
  

  