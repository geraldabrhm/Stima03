import React, { useState } from 'react'
import { Endpoints  } from '../API'

export default function DNATest() {
    const [patientName, setPatientName] = useState("")
    const [diseaseName2, setDiseaseName2] = useState("")
    const [file, setFile] = useState(null)

    function handleSubmit(e) {
        e.preventDefault()
        const formPatient = document.getElementById('form-dis')
        formPatient.reset();
    }
    
    const fileInputHandler = (event) => {
        setFile(event.target.files[0])
    }

    const fileSubmitHandler = async () => {
        var fileInput2 = document.getElementById("patient-file").files[0]
        var dataFetch2 = new FormData()

        dataFetch2.append('PatientName', patientName)
        dataFetch2.append('IDDisease', diseaseName2)
        dataFetch2.append('PatientDNA', fileInput2)
        
        const res = await fetch(Endpoints.patientBM, {
            method: "POST",
            // headers: {
            //     'Content-Type': 'application/json'
            // },
            // header: "origin",
            body: dataFetch2
        }).then(res => {
            return res.json()
        }).catch(error => console.log('ErrorMessage: ', error))

        const lastRes = await fetch(Endpoints.lastResult, {
            method: "GET"
        }).then(res => {
            return res.json()
        }).then(data => {
            console.log(data)
        }).catch(err => {
            console.log('ErrorMessage: (Get) ', err)
        })

        const val = await lastRes.json()
        console.log(val)
    }

    return (
        <section className="container">
            <h1 className="header">DNA Test</h1>
            <form onSubmit={handleSubmit} className="add-form" id="form-dis">
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
                    <input type="file" id="patient-file" onChange={fileInputHandler}/>
                </div>
                <div className="form-control">
                    <label>
                        Disease to predict:
                    </label>
                    <input type="text" id="disease-namee" placeholder="Input the disease id" onChange={e=>setDiseaseName2(e.target.value)}/>
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