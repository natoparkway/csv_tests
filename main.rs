use std::error::Error;
use std::io;
use std::process;

// fn vec_of_vecs_load() -> Result<(), Box<dyn Error>> {
//     let mut data = vec![];
//     let mut rdr = csv::Reader::from_reader(io::stdin());
//     for result in rdr.records() {
//         let record = result?;
//         let mut row = vec![];
//         for value in record.iter() {
//             row.push(value.to_owned());
//         }
//         data.push(row);
//     }
//     Ok(())
// }

fn vec_of_records_load() -> Result<Vec<csv::StringRecord>, Box<dyn Error>> {
    let mut data = vec![];
    let mut rdr = csv::Reader::from_reader(io::stdin());
    for result in rdr.records() {
        let row = result?;
        data.push(row);
    }
    Ok(data)
}

fn make_columnar(data: Vec<csv::StringRecord>) -> Vec<Vec<String>> {
    let mut columar_data = vec![];
    for row in data {
        for (i, elem) in row.iter().enumerate() {
            if columar_data.len() <= i {
                columar_data.push(vec![]);
            }
            columar_data[i].push(elem.to_owned());
        }
    }
    return columar_data;
}

fn main() {
    // cargo build
    // /usr/bin/time -lp cargo run < ../data/importer_contacts50.csv
    if let Ok(data) = vec_of_records_load() {
        let columar_data = make_columnar(data);
        println!("Loaded {} columns", columar_data.len());
        println!("Loaded {} rows", columar_data[0].len())

    }
}
