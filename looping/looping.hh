<?hh

function loop(int $duration): void {
	for ($i = 0; $i < $duration; $i++) {
		echo sprintf("Loop number %d\n", ($i+1));
	}
}

loop(100000);
