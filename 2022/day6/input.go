package day6

var dataInput = "rvnvzvhzzjgjgffclllnhhtltptgptgpttjhttsllmbbphbpbzpbpjjcwjwqwccnrrtvrtrfrwffnqffsggwzzhtzhthqhffmrrzsrrnrtrqqbllhrlrjrvvrvvgdgjgfjjtzznffrfvfggswgssccpcwpccstcsstwssgzssswsgwsgsnsshcsscsffcwcmmhmsmrsrddwhhszsfsjfsfccwssvdssmzzwrzzjfzzvzzfsfsvshschhvlhhvmvpvhhstsdtstgtrrtjjcssgfssnsnfsnscnsslvvqbbhthqtqzqmzmjjfmmbdmmzdmdpmpnnfjnnrwnwbwrrwjrrttfzzlwwfmmnhhqlhqqbnqnqznncmmhfmfftsfszssftstggppfddtztltctqcqjcqcmclcnccdgcggvmvvcsstztbzttsccmgcghcczfzszbzmbmjjggsqgghbhjjhhrrjljmllczlzzqssjfjqfjfmfdfgffczfzwwpfwpwbbftbffplphhfcftctlcclhlnlhnhlhnhqnhhjbjmmtpmmlsmllvhhjghhmzzwpzprrjrzzvrvcvqvssgssnccmgmppfhfwhhczhhprhrffpvvqrqvqccdhdvvvssvdvwdwwzggfvfnvvqvwwzwmzzvttgzzwjjfvvpvffqpqhqrrwccjbjppbzbrbrrphrhgrhhzbbhrrblbjjznnlccsffpcfpcfpfddfwddbffczztzsttzsstddrqdqqchqccrlccfcwwghwhnhbhjbhhzrrgprrftrffqlljlslggzdzvvrgvgcgnnngtnggdcgddhjhchqcqfqwfwjffqppblbddgmdgmddrnddjtdjttgzgpzggwvwjvwwhdwwbvvtftcfttpftpfpzfznnchnhrrdtdbtdbdppsjsgsttptsptssnwntwnnsshqqgnnqnddhjjtqqjhhmghmhmsstrrhmmwvvlssprsrbbswsddbnntmtbmtbmmtmrrffvjffdrffnttbffbccgchghjgjwwshwswdswdsdsrdsdjdvjjgbjbssptpddzbdddqdpdspswssjrjwwcbbvsspnspslpplzlqlccvlvfllntlntnqnrqnqpqlqflfrfsfgsfswwptwwgqggrhghzghhrthtzhzsswhhzffpsffhnnrvnrvvhmmlvvqsslddhcddsllpmpggrhrcrbbmnmjmqmlmpmrrqwqhqrqhrrgrqqsrrpjpdjppzhphthddlzdzhddjvdjdcjcncnvnzzvwzvzgzjggcllvpllgdgqqdvvwnvwvwqqvwvvrbbtqqwhdztqfzzcqrshjzwqnpdsshmpjzwqdptbvfqzmfnbtlgbbsjbqgnblhbbpsfdzvcmpzfwczcnbdndsjzccjcqnrdglwfrvwtnjwpvpvvgwtmnpzhbwnbqwznmdvdrsjnlsfrpcnhlbmlgmrcjbbvhqnvbrmlnfttjllstqnqqnhqrzrhfqjbfbwfhhjzwtwzjmszzhjnjbhrlbnnpfvdmlftjnfsfnvddqfhqqlljrthptvmhbqrdmdcmljwgbrqrjwcrtmvjgqtblnslgbjmdsrrpgfsctrhwlwnszpljhrfnsfpcsgczzltgsztclhgcqljrcmbbpdlztncnrnrmgrttplcnldfqddqhznmcczbmwvsztmwmcqnrzmlmqchnhnhrrhhfntjzqcbnttspptqwvphlbtpfcbhdqzbhsbvhmsqbsdntwpcrnzvmbgqsgttbqhhblfjpmvfcrzhfnzwrzbzgsdfqndzzhfdnrvsnvfbptthjnhgljhrvwwrlbnfpvvjdjchcgbhfrqvszhrhqvtzplwsptvdhqwlzhcjpclmmrlccvgvtgsfpnjhqhrbqglznpdhmsqwwsbmhmlsmmvvghsmplqjchrfctltmnqnddzqjfpljwljbdqjcqdqzwsbcclqsmsmlstvljwwtfmpnhqzqfjghjfchjccqrchsvngvrnwfwttsnvrdlfvwfptsjcpslvvpmjclfcpljqjszptsgsmntzrdjbgrzmgvzddqrlsndjzzqbznqnphbnwfhtjjlwjpsvffrdrbsbttpgrvmqrdndqvlgzcpfbttvqdgvrmtvfclcbcwllthmdzjcflwpnrsqzrjdzbvqgzsqvjjpjjpnjtcjqhcfbjdqndlcwzhcbjtgtlvtdwctdnqcbcgsmrcrmwjntdwbjdbzmshbvlspjfdbvmlrbdzlmlggthvphnrqlrcdsqpsgqcmpgmgdzvdqlmcldvztpsbmpwjgjfhswplwrvwpbwbsgsvlhdvmpzwnnbwjwmshwcnqfqjdpchfjjcbdnslqchhnznpqpnnctznbtccclgjhmnngdjlmqnzpsdptqqcrblmrlnnpgvrfrtmbjnmspfrbwpclhgtbsghndrjfbggsplvcjnhjjbqzsfdpfnvchzjgbhdqgddfgddvzdcrjlntnsmscqwmpptqgbnvtsvpjvmhcfpbfrpzqbpfhlbjrmbmvdvvnvfqsndglvhfrqcsbsbbprscrbfthzcwcdrprqrwjzwrpblfllpwzlhmqvltjgpcjzpzwbltgwsrgrrhzcqfpvhcprdhnzcfphqrwcvtpcbppjwzmmwjhbvwbblnbwvcqzvlfzjhgmlnhlhrbsplfctggfbhbgwpncznvtmdtqqmjsvsnrlqswzvflrfsncgpdcndlwrfrqwqnqtjmsphwsgzhdjpnsdgrbrfhbfdrntwvgvbwnvwnrmdbhqgrglbfwprflnrljrwsgwtpgtmfhvvghtzndvwlzjchhmlwcncmpvrslrglzjcnhfdqhcrljhgbzpssvdnmfwzstmvrztgpsscfswltnbwrrtcnvbswmmjbmnnnvqwjzhprfnvlbvzzdvbwlwchrvqnwwpbnttbhfdvjjvzsznhczjcncmcrmwtrlsvbwpsrcwqdvgcfjsbqnwmjmtcgpnmcbfbcqhzrjbtlpvwzhjqqprbdnbgzfwlprlcspwjwnfftldqzbcgqnjtglvbpqffdvjbpslqcdzwdnmncvcwfshdhsmssttqfrsbnjgmhfqzlgrbpdfqtfdwslsgphfzzgbzjssfbgnwztzmczplqwjmhtlflpvqqqmrvlllhngtfgsvbbnvhzqbcgpmnlsmpwqwgqfjpzplzjhrslwzrsrgjgpppjlnhrnggcdzvspsztnschnqftgffbtvrpzndzpqmtsfmwgnrvpmtgpvnmqfmwgcvlznwqnjnjwpgnqfjwtdhhmztlvlcvrlzmlpmjdvdnzwbfsshfsvbbhqsmphjhqtlnvlsmvwlqvfqpnjnzlgmdvwzzrllmcgwtqdwphhbwmlrqhqrfmdvtqswvsllqvwmfbbpllsbjbvgsmrcgqvfgsnszfrdlcjbtfgcbclhmzlmqlnhmslcmgrcvjjlbpsjjcznzqwmztcgdgbmrlgwzjzzjrpndfrdzztzzgmsrcnwrvqrdcczrbhdpfjwvqsmbrcvllvjrrbrsqnzldltdgzscjrssvdhzhnvltpgfdfcvfbtqmphdzhpzgjhjbwsmdlbqqgcrhjrwhgvfgdllmmlnpsbtvdrrzmltfwgrcsfrrsdbdmjtrwhnlgcrgjgmbzvzvqbflvbrsqcssjgsvpjmlhtcggzwbvddmwwtfdrlltjqcpwnthczzlzdszvtmrmhbpcgstvrsnmbdcmjzsvnncmgmlnrzzhfvmblgptwwwbrmtcczjwcqmvdrsvfjgqqvghnhbntqcdrppfmvdzbcjvztrnpdhmdpmnvsnzzldvdfqbdrwqqqqsmswthwpjwdrflszbspqhpfwztjjcbdsrftdsrsfdltnfztcslmbsghgrtcscrfmptqplwpmtqdzthgfjhbqdnsffrpjmgczmrlfvzcjttwtmtqfbtlqrttvjdwhfgcgcclrlswmzhzbfhjrggnhwtnffnqqcvldlttvvgrbcqbmqzvtflfmdblhdbzphrqtbshvp"