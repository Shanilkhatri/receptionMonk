import React from "react";
import Avatar from "react-avatar-edit";
import { useState } from "react";

function Foravatar() {
  const [preview, setPreview] = useState(null);
  function onClose() {
    setPreview(null);
  }
  function onCrop(pv: any) {
    setPreview(pv);
  }
  function onBeforeFileLoad(elem: any) {
    if (elem.target.files[0].size > 5000000) {
      alert("File is too big!");
      elem.target.value = "";
    }
  }
  return (
    <div className="m-4">
      (image size up to 5MB, accepts only png and jpeg formats )
      <Avatar
        width={150}
        height={150}
        onCrop={onCrop}
        onClose={onClose}
        onBeforeFileLoad={onBeforeFileLoad}
        src={null}
      />
      <br />
      {preview && (
        <>
          <img src={preview} alt="Preview" />
          <a href={preview} download="avatar">
            Download image
          </a>
        </>
      )}
    </div>
  );
}
export default Foravatar;
