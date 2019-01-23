function computeLPSArray(pattern) {
    // length of the previous longest prefix-suffix
    var len = 0;
    var lps = [0];
    var i = 1;
    const m = pattern.length;

    while(i < m) {
        if(pattern[i] === pattern[len]) {
            len++;
            lps[i] = len; // store len in lps array
            i++;
        } else {
            if(len !== 0) {
                len = lps[len - 1];
            } else {
                lps[i] = 0;
                i++;
            }
        }
    }

    console.log(lps);
}

computeLPSArray("AABAACAABAA");