function computeLPSArray(pattern) {
    // length of the previous longest prefix-suffix
    var j = 0;
    var lps = [0];
    var i = 1;
    const m = pattern.length;

    while(i < m) {
        if(pattern[i] === pattern[j]) {
            j++;
            lps[i] = j; // store len in lps array
            i++;
        } else {
            if(j !== 0) {
                j = lps[j - 1];
            } else {
                lps[i] = 0;
                i++;
            }
        }
    }
    console.log("LPS: ", lps);
    return lps;
}

function KMPSearch(pattern, text) 
{
    console.log(pattern);
    console.log(text);
    var patLength = pattern.length;
    var textLength = text.length;

    var lps = computeLPSArray(pattern);
    var textIncrement = 0;
    var patternIncrement = 0;

    while (textIncrement < textLength) { 
        // console.log("patternIncrement: ", patternIncrement);
        // console.log("textIncrement: ", textIncrement);
        if (pattern[patternIncrement] === text[textIncrement]) { 
            patternIncrement++;
            textIncrement++;
        }
        if (patternIncrement === patLength) {
            console.log("Found pattern at index ", textIncrement - patternIncrement);
            patternIncrement = lps[patternIncrement - 1]; // Here is the optimization.  It is reverting back to the lps array.  If there is a prefix thatis also a suffix, it can skip to the suffix and check that
        } else if (textIncrement < textLength && pattern[patternIncrement] !== text[textIncrement]) {  // mismatch after j matches
            // Do not match lps[0..lps[j-1]] characters,
            // they will match anyway
            if ([patternIncrement] != 0)
                patternIncrement = lps[patternIncrement - 1];
            else
                textIncrement = textIncrement + 1;
        }
    }
}

//computeLPSArray("AABAACAABAA");
const text = "ABABDABACDABABCABAB";
const pattern = "ABABCABAB";
KMPSearch(pattern, text);