import { Endpoints  } from '../API'

export default function DNATestResult() {
    function handleSubmit(e) {
        e.preventDefault();
    }

    const fileSubmitHandler = async (event) => {
        // ! Server, Endpoints belum dibuat
        // const res = await fetch(Endpoints.dna_test_result, {
        //     method: "GET", // GET, POST, PUT, PATCH
        //     headers: {
        //         'Content-Type': 'application/json'
        //     },
        // }).then(res => {
        //     return res.json()
        // }).then(data => console.log(data))
        // .catch(error => console.log('ErrorMessage'))

        const date_val = document.getElementById("dna_result_date").value
    }

    const prevResult = () => {
        
    }

    const nextResult = () => {

    }



    return (
        <section className="container2">
            <h1>DNA Test Result</h1>
            <form onSubmit={handleSubmit} className="add-form">
                <div className="form-control">
                    <label>
                        Date test:
                    </label>
                    <input type="date" id="dna_result_date"/>
                </div>
                <div className="form-control">
                    <label>
                        Disease name:
                    </label>
                    <select>
                        <option>Penyakit 1</option>
                        <option>Penyakit 2</option>
                        <option>Penyakit 3</option>
                        <option>Penyakit 4</option>
                    </select>
                </div>
                <input type='submit' value='Find Result' className = 'btn' onClick={fileSubmitHandler}/>
            </form>
            <ul>
                <li>A</li>
                <li>B</li>
                <li>C</li>
                <li>D</li>
                <li>E</li>
            </ul>
            <input type='submit' value='Previous' className = 'btn-prev' onClick={nextResult}/>
            <input type='submit' value='Next' className = 'btn-next' onClick={nextResult}/>
            
        </section>
    )
}
// TODO get list of DiseaseName from server
// TODO get list of DNATestResult from server