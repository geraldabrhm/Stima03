import React, { useState } from 'react'
import { Endpoints  } from '../API'

export default function DNATest() {
    const [patientName, setPatientName] = useState("")
    const [file, setFile] = useState(null)

    function handleSubmit(e) {
        e.preventDefault()
    }
    
    const fileInputHandler = (event) => {
        setFile(event.target.files[0])
    }

    const fileSubmitHandler = async (event) => {
        // ! Server, Endpoints belum dibuat
        // TODO Handle file dan select option juga
        // const res = await fetch(Endpoints.dna_test, {
        //     method: "PATCH", // GET, POST, PUT, PATCH
        //     headers: {
        //         'Content-Type': 'application/json'
        //     },
        //     body: JSON.stringify({
        //         patientName
        //     })
        // }).then(res => {
        //     return res.json()
        // }).then(data => console.log(data))
        // .catch(error => console.log('ErrorMessage'))
    }

    return (
        <section className="container">
            <h1 className="header">DNA Test</h1>
            <form onSubmit={handleSubmit} className="add-form">
                <div className="form-control">
                    <label>
                        Patient name:
                    </label>
                    <input type="text" placeholder="Input the patient name" onChange={e=>setPatientName(e.target.value)}/>
                </div>
                <div className="form-control">
                    <label>
                        DNA sequence:
                    </label>
                    <input type="file" onChange={fileInputHandler}/>
                </div>
                <div className="form-control">
                    <label>
                        Disease to predict:
                    </label>
                    <select>
                        <option>Penyakit 1</option>
                        <option>Penyakit 2</option>
                        <option>Penyakit 3</option>
                        <option>Penyakit 4</option>
                    </select>
                </div>
                <div className="form-control">
                    <label>
                        Algoritma pattern-matching:
                    </label>
                    <select>
                        <option>Knuth-Morris-Pratt</option>
                        <option>Boyer-Moore</option>
                    </select>
                </div>

                <input type='submit' value='Submit Test' className = 'btn' onClick={fileSubmitHandler}/>
            </form>
            <h2>Test Result</h2>
            <p id="resultTest">Result</p>
        </section>
    )
}
// TODO get list of DiseaseName from server