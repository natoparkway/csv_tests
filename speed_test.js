const Papa = require("papaparse");
const fs = require('fs');
const {performance} = require('perf_hooks');
const _ = require("lodash");




function parsePapa(fname, startTime) {
	const file = fs.createReadStream(fname);
	Papa.parse(file, {
		complete: function(results) {
			console.log(
				fname, ':', results.data.length, performance.now() - startTime
			)
			// printMemoryUsage('cb')
		}
	});
} 

function test_fn_speed(fn, n) {
	// Returns fn speed in miliseconds
	var durations = []
	for (let i = 0; i < n; i++) {
		const startTime = performance.now()
		fn()
		durations.push(performance.now() - startTime)
	}

	const totalDuration = durations.reduce((a, b) => a + b, 0)
	return totalDuration / durations.length  

}



function printMemoryUsage(label) {
	console.log(label)
	const used = process.memoryUsage();
	for (let key in used) {
  		console.log(`${key} ${Math.round(used[key] / 1024 / 1024 * 100) / 100} MB`);
	}
}

function main() {
	// 'data/importer_contacts10k.csv',
	// 'data/importer_contacts_100k.csv',
	// 'data/importer_contacts1M.csv',
	// 'data/importer_contacts10M.csv'
	const fname = 'data/importer_contacts10M.csv'
	const file = fs.createReadStream(fname);
	Papa.parse(file, {
		complete: function(results) {
			console.log(results.data.length)
		}
	});


}

main()

// const options = {/* options */};

// const dataStream = fs.createReadStream("data/importer_contacts50.csv");
// const parseStream = papa.parse(papa.NODE_STREAM_INPUT, options);

// dataStream.pipe(parseStream);

// let data = [];
// parseStream.on("data", chunk => {
//     data.push(chunk);
// });

// parseStream.on("finish", () => {
//     console.log(data);
//     console.log(data.length);
// });