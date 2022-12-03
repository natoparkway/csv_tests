import sys
import csv
import pandas

# process = psutil.Process(os.getpid())
# print(process.memory_info().rss)  # in bytes 
# print(resource.getrusage(resource.RUSAGE_SELF).ru_maxrss)

def read_csv_with_csv(fname):
	with open(fname) as f:
		reader = csv.DictReader(f, delimiter=',', quotechar='"')
		data = [row for row in reader]
	return data

def read_csv_with_pandas(fname):
	df = pandas.read_csv(fname, delimiter=',', quotechar='"')
	return df


# def test_fn_speed(fn, iterations=1):
# 	times_elapsed = []
# 	for i in range(iterations):
# 		beginning = time.time()
# 		fn()
# 		times_elapsed.append(time.time() - beginning)
# 	return round(sum(times_elapsed) / len(times_elapsed), 4)


def main():
	# ./harness.sh "python3 speed_test.py data/importer_contacts1M.csv pandas"
	fname = sys.argv[1]
	fn = sys.argv[2]

	csv_fns = {
		'csv': read_csv_with_csv,
		'pandas': read_csv_with_pandas
	}

	print(fname)
	result = csv_fns[fn](fname)
	# if fn == 'pandas':
		# print(result.info())




if __name__ == '__main__':
	main()