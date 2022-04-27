import React, { useState } from 'react'
import { Endpoints  } from '../API'

export default function InputDisease() {
    const [diseaseName, setDiseaseName] = useState("")
    const [file, setFile] = useState(null)

    function handleSubmit(e) {
        e.preventDefault();
        const formDisease = document.getElementById('disease-form')
        formDisease.reset();
    }

    const fileInputHandler = (event) => {
        setFile(event.target.files[0])
    }

    const fileSubmitHandler = async (event) => {
        var fileInput = document.getElementById("disease-file").files[0]
        var dataFetch = new FormData()

        dataFetch.append('DiseaseName', diseaseName)
        dataFetch.append('DNASequence', fileInput)

        const res = await fetch(Endpoints.disease, {
            method: "POST", // GET, POST, PUT, PATCH
            // headers: {
            //     'Content-Type': 'application/json'
            // },
            body: dataFetch
        }).then(res => {
            return res.json()
        }).then(data => console.log(data))
        .catch(error => console.log('ErrorMessage: ', error))
    }

    return (
        <section className="container">
            <h1 className="header">Input Disease</h1>
            <form onSubmit={handleSubmit} className="add-form" id="disease-form">
                <div className="form-control">
                    <label>
                        Disease name:
                    </label>
                    <input type="text" id="disease-text" placeholder="Input the disease name" onChange={e=>setDiseaseName(e.target.value)}/>
                    {/* <input type="text" placeholder="Input the disease name" onChange={e=>setDiseaseName(e.target.value)}/> */}
                </div>
                <div className="form-control">
                    <label>
                        DNA sequence:
                    </label>
                    <input type="file" id="disease-file" onChange={fileInputHandler}/>
                </div>
                <input type='submit' value='Submit Disease' className = 'btn' onClick={fileSubmitHandler}/>
                {/* <input type='submit' value='Submit Disease' className = 'btn' onChange={e=>setDiseaseName(e.target.value)}/> */}
            </form>
        </section>
    )
}