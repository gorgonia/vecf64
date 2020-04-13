void sum(double a[], int len, double* retVal) {
	double acc = 0.0;
	for (int i = 0 ; i < len; i++){
		acc += a[i];
	}
	*retVal = acc;
}
