import React, { useState } from 'react'
import { Endpoints  } from '../API'

export default function DNATest() {
    const [patientName, setPatientName] = useState("")
    const [diseaseName, setDiseaseName] = useState("")
    const [file, setFile] = useState(null)

    function handleSubmit(e) {
        e.preventDefault()
    }
    
    const fileInputHandler = (event) => {
        setFile(event.target.files[0])
    }

    const fileSubmitHandler = async () => {
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
        var head = {
            'Access-Control-Allow-Origin': 'http://localhost:8080/api/disease/',
            'Access-Control-Allow-Credentials': 'true',
        }

        const res = await fetch("http://localhost:8080/api/disease/", {
            method: "GET",
            // mode: "no-cors",
        }).then(res => {
            console.log("Ini res")
            console.log(res)
            console.log("Ini res.json()")
            console.log(res.json())
            res.json()
        }).then(data => {
            console.log("Ini data")
            console.log(data)
        }).catch(err => {
            console.log("Error bingung: ", err)
        })
        document.getElementById("patient-name").value = res
    }

    return (
        <section className="container">
            <h1 className="header">DNA Test</h1>
            <form onSubmit={handleSubmit} className="add-form">
                <div className="form-control">
                    <label>
                        Patient name:
                    </label>
                    <input type="text" id="patient-name" placeholder="Input the patient name" onChange={e=>setPatientName(e.target.value)}/>
                </div>
                <div className="form-control">
                    <label>
                        DNA sequence:
                    </label>
                    <input type="file" id="patient-dna-file" onChange={fileInputHandler}/>
                </div>
                <div className="form-control">
                    <label>
                        Disease to predict:
                    </label>
                    <input type="text" id="disease-name" placeholder="Input the disease name" onChange={e=>setDiseaseName(e.target.value)}/>
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