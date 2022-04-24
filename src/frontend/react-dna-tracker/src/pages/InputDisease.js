import React, { useState } from 'react'
// import axios from 'axios'

export default function InputDisease() {
    const [diseaseName, setDiseaseName] = useState("")
    const [file, setFile] = useState(null)

    function handleSubmit(e) {
        e.preventDefault();
    }

    const fileInputHandler = (event) => {
        setFile(event.target.files[0])
    }

    const fileSubmitHandler = (event) => {
        // axios or sequelize
    }

    return (
        <section className="container">
            <h1 className="header">Input Disease</h1>
            <form onSubmit={handleSubmit} className="add-form">
                <div className="form-control">
                    <label>
                        Disease name:
                    </label>
                    <input type="text" placeholder="Input the disease name" onChange={e=>setDiseaseName(e.target.value)}/>
                    {/* <input type="text" placeholder="Input the disease name" onChange={e=>setDiseaseName(e.target.value)}/> */}
                </div>
                <div className="form-control">
                    <label>
                        DNA sequence:
                    </label>
                    <input type="file" onChange={fileInputHandler}/>
                </div>
                <input type='submit' value='Submit Disease' className = 'btn' onChange={fileSubmitHandler}/>
                {/* <input type='submit' value='Submit Disease' className = 'btn' onChange={e=>setDiseaseName(e.target.value)}/> */}
            </form>
        </section>
    )
}