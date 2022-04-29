// const heroku = "https://cors-anywhere.herokuapp.com/"
export const apiURl = "http://localhost:8080" /*.concat("/proxy")*/

export const Endpoints = {
  disease: `${apiURl}/api/disease`,
  patientKMP: `${apiURl}/api/patient/kmp`,
  patientBM: `${apiURl}/api/patient/bm`,
  lastResult: `${apiURl}/api/prediction-last`,
  prediction: `${apiURl}/api/prediction`
  // dna_test: `${apiURl}/api/dna-test`,
  // dna_test_result: `${apiURl}/api/dna-test-result`,
}