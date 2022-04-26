import React, { useState } from 'react'
import { Endpoints  } from '../API'

export default function InputDisease() {
    const [diseaseName, setDiseaseName] = useState("")
    const [file, setFile] = useState(null)

    function handleSubmit(e) {
        e.preventDefault();
    }

    const fileInputHandler = (event) => {
        setFile(event.target.files[0])
    }

    const fileSubmitHandler = async (event) => {
        // ! Server, Endpoints belum dibuat
        // TODO Handle file juga
        // const res = await fetch(Endpoints.dna_test, {
        //     method: "PATCH", // GET, POST, PUT, PATCH
        //     headers: {
        //         'Content-Type': 'application/json'
        //     },
        //     body: JSON.stringify({
        //         diseaseName
        //     })
        // }).then(res => {
        //     return res.json()
        // }).then(data => console.log(data))
        // .catch(error => console.log('ErrorMessage'))
        let a = document.getElementById("disease-file").files[0]
        console.log(a)
    }

    return (
        <section className="container">
            <h1 className="header">Input Disease</h1>
            <form onSubmit={handleSubmit} className="add-form">
                <div className="form-control">
                    <label>
                        Disease name:
                    </label>
                    <input type="text" id="disease-text"placeholder="Input the disease name" onChange={e=>setDiseaseName(e.target.value)}/>
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