class tree:
    def __init__(self, data):
        self.sequence = data
        self.children = []
    def add_child(self, obj):
        self.children.append(obj)

def smallestMotif(motifs):
    smallestMotif = motifs[0]
    for motif in motifs:
        if len(motif) < len(smallestMotif):
            smallestMotif = motif
    return smallestMotif
def subsequencesI(motif):
    if(len(motif)<=1):
        return 
    else:
        print("the father")
        print(motif)
        print("the father")
        for i in range(len(motif)):
            print(motif[:i] + motif[i + 1:])
            subsequencesI(motif[:i] + motif[i + 1:])
def subsequences(motifs):
 motif = smallestMotif(motifs)
 subsequencesI(motif)

subsequences(["ATGjsfdj", "ATGksdgk", "ATGjasdjg", "ryad"])

 