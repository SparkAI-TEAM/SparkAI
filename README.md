## SparkAI Project

Official Golang implementation of the Spark AI  Chain protocol.

This project was forked from the  go-ethereum. And  it was   changed  to work with Spark AI Chain.

And We are working on.

Automated builds are available for stable releases and the unstable master branch. 

Spark AI  frogenet V1.0   will  landing  at  June.2025

Spark Swap & crosschain bridge is developing in progress.

## Building the source

Information abount  Spark AI  & POY


Spark AI will be the gas of our chain.
Poy will be a contract  token.

Spark  AI  totalSupply : 190,000,000
POY totalSupply: 100,000,000

## Building the source

Building `spark` requires Go (version 1.17 or later)  . You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
cd cmd
cd gspark
go build
```



## Executables

The gspark programe is Our main Forgnet Chain CLI client. It is the entry point into the spark network , capable of running as a full node (default), archive node (retaining all historical state) or a light node (retrieving data live).  

You can run it as a node to sync and query chain-data for the BpFlac Chain.

Forgenet connect only import the belivable nodes.To do our  AI transformer work to mine out.


Some cpu power of  spark mine:

#  CPU HASHRATE  RUNSAMPLES

1 AMD Eng Sample: 100-000000053-04_32/20_N 101803.97 22

2 AMD EPYC 7763 64-Core Processor 98527.99 3

3 AMD EPYC 7742 64-Core Processor 95400.73 16

4 AMD EPYC 7H12 64-Core Processor 93554.12 1

5 AMD EPYC 7702 64-Core Processor 85965.3 18

6 AMD EPYC 7713 64-Core Processor 80677.04 7

7 AMD EPYC 7K62 48-Core Processor 79678.74 10

8 AMD EPYC 7552 48-Core Processor 71423.47 2

9 AMD Ryzen Threadripper 3990X 64-Core Processor 65116.88 86

10 AMD EPYC 7V12 64-Core Processor 60262.75 8

11 Intel® Xeon® Platinum 8153 CPU @ 2.00GHz 54570.26 1

12 AMD Ryzen Threadripper PRO 3995WX 64-Cores 54259.36 11

13 AMD EPYC 7542 32-Core Processor 53980.82 4

14 AMD EPYC 7452 32-Core Processor 53763.44 25

15 Intel® Xeon® Platinum 8268 CPU @ 2.90GHz 46814.29 1

16 AMD Eng Sample: ZS1406E2VJUG5_22/14_N 45435.96 18

17 AMD EPYC 7402 24-Core Processor 45398.83 9

18 AMD Eng Sample: 2S1404E2VJUG5_20/14_N 44008.27 11

19 AMD EPYC 7F72 24-Core Processor 43457.48 1

20 AMD EPYC 7J13 64-Core Processor 43085.06 4

21 AMD EPYC 7413 24-Core Processor 42518.81 8

22 AMD Ryzen Threadripper 3970X 32-Core Processor 42365.7 110

23 AMD EPYC 7702P 64-Core Processor 40321.93 10

24 AMD EPYC 7R32 38725.17 17

25 AMD Eng Sample: 100-000000053-05_32/20_Y 38529.71 2

26 Genuine Intel® CPU 0000 @ 2.10GHz 38172.46 6

27 Intel® Xeon® Gold 6150 CPU @ 2.70GHz 35432.97 3

28 AMD EPYC 7601 32-Core Processor 34973.11 23

29 AMD EPYC 7501 32-Core Processor 33179.6 2

30 AMD EPYC 7551 32-Core Processor 32390.76 6

31 AMD Ryzen Threadripper 3960X 24-Core Processor 30598.82 43

32 AMD Eng Sample: ZS1711E3VIVG5_24/17_N 30256.91 18

33 Intel® Xeon® Platinum 8358 CPU @ 2.60GHz 30250.78 2

34 Intel® Xeon® CPU E7-8890 v3 @ 2.50GHz 29713.62 2

35 AMD Eng Sample: 100-000000057-04_31/22_N 29079.91 8

36 Intel® Xeon® Platinum 8269CY CPU @ 2.50GHz 28971.23 3

37 AMD EPYC 7502P 32-Core Processor 28413.13 13

38 AMD EPYC 7302 16-Core Processor 28312.57 9

39 Intel® Xeon® Platinum 8275CL CPU @ 3.00GHz 27414.01 13

40 Intel® Xeon® Platinum 8373C CPU @ 2.60GHz 27222.59 1

41 Intel® Xeon® Platinum 8180 CPU @ 2.50GHz 26195.84 4

42 Intel® Xeon® Gold 6230R CPU @ 2.10GHz 24886.77 3

43 Intel® Xeon® CPU E7-8880 v3 @ 2.30GHz 24748.8 8

44 Intel® Xeon® Platinum 8272CL CPU @ 2.60GHz 24580.29 10

45 Intel® Xeon® Gold 6258R CPU @ 2.70GHz 24576.06 9

46 Intel® Xeon® Gold 6248R CPU @ 3.00GHz 24548.31 3

47 AMD EPYC 7571 23213.71 5

48 Intel® Xeon® W-3175X CPU @ 3.10GHz 22938.2 7

49 AMD EPYC 7571 32-Core Processor 22870.73 1

50 AMD Ryzen 9 5950X 16-Core Processor 22763.49 1362

51 Intel® Xeon® Platinum 8124M CPU @ 3.00GHz 22710.24 5

52 AMD Ryzen 9 3950X 16-Core Processor 22520.49 1114

53 AMD Eng Sample: 100-000000054-04_32/24_N 22449.71 1

54 AMD Eng Sample: 2S1905A4VIHF4_30/19_N 22148.88 2

55 AMD EPYC 7D12 32-Core Processor 22069.21 2

56 Intel® Xeon® Gold 6238R CPU @ 2.20GHz 21955.82 3

57 Intel® Xeon® Platinum 8160 CPU @ 2.10GHz 21635.66 1

58 AMD Ryzen Threadripper 2990WX 32-Core Processor 21454.62 31

59 Intel® Xeon® Gold 6238 CPU @ 2.10GHz 21283.84 3

60 AMD EPYC 7402P 24-Core Processor 20633.49 4

61 Intel® Xeon® Gold 6246R CPU @ 3.40GHz 20561.75 3

62 Intel® Xeon® CPU E7-8891 v2 @ 3.20GHz 19921.11 10

63 AMD EPYC 7451 24-Core Processor 19875.97 3

64 AMD EPYC 7352 24-Core Processor 19774.57 3

65 Intel® Xeon® Gold 6248 CPU @ 2.50GHz 19135.46 1

66 Intel® Xeon® Platinum 8280L CPU @ 2.70GHz 18986.72 2

67 Intel® Xeon® Gold 6148 CPU @ 2.40GHz 18726.94 9

68 Intel® Xeon® Gold 6242 CPU @ 2.80GHz 18585.63 6

69 Intel® Xeon® CPU E5-2696 v4 @ 2.20GHz 18376.61 2

70 Intel® Xeon® Platinum 8175M CPU @ 2.50GHz 18373.1 3

71 Intel® Xeon® Gold 6240R CPU @ 2.40GHz 18367.83 1

72 Intel® Xeon® Platinum 8259CL CPU @ 2.50GHz 18253.34 12

73 Intel® Xeon® Gold 6226R CPU @ 2.90GHz 18094.01 5

74 AMD Ryzen Threadripper PRO 3955WX 16-Cores 17825.95 6

75 Intel® Xeon® Gold 6139 CPU @ 2.30GHz 17759.78 2

76 AMD Ryzen 9 3900X 12-Core Processor 17596.34 1628

77 Intel® Xeon® Gold 6130 CPU @ 2.10GHz 17365.33 7

78 AMD Ryzen 9 5900X 12-Core Processor 17310.62 579

79 Intel® Xeon® CPU E5-2699 v4 @ 2.20GHz 17309.12 6

80 Intel® Xeon® CPU E7-4830 v3 @ 2.10GHz 17204.6 1

81 Intel® Xeon® Gold 5218R CPU @ 2.10GHz 17143.25 7

82 Intel® Xeon® Gold 5218 CPU @ 2.30GHz 17031.71 15

83 AMD Ryzen 9 3900XT 12-Core Processor 16996.4 110

84 AMD Ryzen Threadripper PRO 3975WX 32-Cores 16775.43 2

85 Intel® Xeon® Silver 4216 CPU @ 2.10GHz 16665.83 8

86 Intel® Xeon® CPU E5-2686 v3 @ 2.00GHz 16567.54 3

87 Intel® Xeon® Platinum 8276 CPU @ 2.20GHz 16414.11 1

88 Intel® Xeon® Gold 6230 CPU @ 2.10GHz 16402.05 1

89 AMD EPYC 7551P 32-Core Processor 16386.75 1

90 Intel® Xeon® CPU E5-2697R v4 @ 2.30GHz 16052.65 4

91 Intel® Xeon® CPU E5-2697 v4 @ 2.30GHz 15642.11 2

92 Intel® Xeon® W-3275M CPU @ 2.50GHz 15357.23 2

93 Intel® Xeon® Gold 6136 CPU @ 3.00GHz 15329.66 1

94 Intel® Xeon® CPU E5-2696 v3 @ 2.30GHz 15325.91 7

95 Intel® Xeon® CPU E5-2699 v3 @ 2.30GHz 15304.33 13

96 AMD Ryzen Threadripper 2970WX 24-Core Processor 15179.34 10

97 Intel® Xeon® CPU E5-2695 v4 @ 2.10GHz 15038.72 6

98 Intel® Xeon® Gold 6132 CPU @ 2.60GHz 14926.26 1

99 AMD EPYC 7302P 16-Core Processor 14794.43 3

100 Intel® Xeon® CPU E5-2698 v3 @ 2.30GHz 14676.31 13

101 Intel® Xeon® Gold 6246 CPU @ 3.30GHz 14666.84 3

102 Intel® Xeon® Gold 6146 CPU @ 3.20GHz 14550.75 1

103 Intel® Xeon® CPU E5-2680 v4 @ 2.40GHz 14539.53 66

104 Intel® Xeon® Gold 5120 CPU @ 2.20GHz 14505.37 3

105 Intel® Xeon® Gold 6154 CPU @ 3.00GHz 14365.96 1

106 AMD Ryzen 9 3900 12-Core Processor 14259.84 27

107 AMD EPYC 7282 16-Core Processor 14218.68 9

108 AMD EPYC 7401P 24-Core Processor 14178.77 2

109 Intel® Xeon® CPU E5-2683 v4 @ 2.10GHz 14038.83 10

110 Intel® Xeon® CPU E5-2690 v4 @ 2.60GHz 13980.01 13

111 Intel® Core™ i9-7980XE CPU @ 2.60GHz 13970.58 15

112 Intel® Xeon® Gold 6252N CPU @ 2.30GHz 13884.88 1

113 Intel® Xeon® Gold 6126 CPU @ 2.60GHz 13872.32 7

114 Intel® Xeon® CPU E5-2682 v4 @ 2.50GHz 13845.24 4

115 ARM Neoverse-N1 13753.46 8

116 Intel® Xeon® CPU E5-2697 v3 @ 2.60GHz 13730.23 18

117 Intel® Xeon® CPU E5-4620 v2 @ 2.60GHz 13600.17 4

118 AMD Ryzen Threadripper PRO 3945WX 12-Cores 13511.69 2

119 Intel® Xeon® Gold 5220R CPU @ 2.20GHz 13502.38 2

120 Intel® Xeon® Gold 6226 CPU @ 2.70GHz 13438.13 6

121 Intel® Xeon® Platinum 8168 CPU @ 2.70GHz 13333.87 6

122 Intel® Xeon® CPU E5-4627 v2 @ 3.30GHz 13199.82 2

123 Intel® Xeon® Gold 6212U CPU @ 2.40GHz 13195.4 2

124 Intel® Xeon® CPU E5-4657L v2 @ 2.40GHz 13151.49 1

125 Intel® Xeon® CPU E5-2678 v3 @ 2.50GHz 12938.45 139

126 Intel® Xeon® CPU E5-2673 v4 @ 2.30GHz 12853.8 16

127 Intel® Xeon® Gold 5220 CPU @ 2.20GHz 12807.3 2

128 Intel® Xeon® Platinum 8176 CPU @ 2.10GHz 12601.12 1

129 Intel® Xeon® CPU E5-2690 v3 @ 2.60GHz 12561.08 14

130 Intel® Xeon® Gold 6252 CPU @ 2.10GHz 12553.04 2

131 Intel® Xeon® Gold 6140 CPU @ 2.30GHz 12536.2 11

132 Intel® Xeon® Silver 4214 CPU @ 2.20GHz 12450.97 7

133 AMD Ryzen 7 5800X 8-Core Processor 12445.4 409

134 Intel® Xeon® Gold 5118 CPU @ 2.30GHz 12420.2 18

135 AMD EPYC 7351P 16-Core Processor 12259.56 3

136 Intel® Xeon® Platinum 8259L CPU @ 2.50GHz 12249.44 2

137 AMD Opteron™ Processor 6386 SE 12230.62 5

138 Intel® Xeon® Gold 6256 CPU @ 3.60GHz 12072.19 1

139 Genuine Intel® CPU 0000 @ 2.20GHz 12051.68 2

140 AMD Ryzen 7 3700X 8-Core Processor 12032.54 780

141 Intel® Xeon® Gold 6240 CPU @ 2.60GHz 11716.05 2

142 Intel® Xeon® CPU E5-4650 0 @ 2.70GHz 11657.46 3

143 AMD Ryzen Threadripper 1950X 16-Core Processor 11621.56 45

144 HiSilicon Kunpeng-920 11574.85 2

145 Intel® Xeon® CPU E5-2683 v3 @ 2.00GHz 11486.6 12

146 Intel® Core™ i9-10980XE CPU @ 3.00GHz 11429.29 8

147 AMD Ryzen 7 3800X 8-Core Processor 11389.65 180

148 Intel® Xeon® CPU E5-4640 0 @ 2.40GHz 11251.5 14

149 Intel® Xeon® CPU E5-2680 v3 @ 2.50GHz 11243.91 42

150 AMD Opteron™ Processor 6380 11204.98 5

151 Intel® Xeon® Gold 6250 CPU @ 3.90GHz 11201.72 1

152 Intel® Xeon® CPU E7- 4870 @ 2.40GHz 11172.93 7

153 Intel® Xeon® Gold 5215 CPU @ 2.50GHz 11015.04 2

154 Intel® Xeon® Gold 6244 CPU @ 3.60GHz 10942.6 1

155 Intel® Xeon® Gold 6210U CPU @ 2.50GHz 10933.62 4

156 AMD Ryzen 7 PRO 3700 8-Core Processor 10915.84 43

157 Intel® Xeon® Platinum 8252C CPU @ 3.80GHz 10811.4 2

158 Intel® Xeon® Gold 6234 CPU @ 3.30GHz 10642.6 2

159 Genuine Intel® CPU 0000 @ 2.00GHz 10625.75 6

160 Intel® Xeon® CPU E7- 4860 @ 2.27GHz 10585.26 2

161 Intel® Xeon® CPU E5-2696 v2 @ 2.50GHz 10408.1 18

162 Intel® Xeon® CPU E5-2686 v4 @ 2.30GHz 10396.31 10

163 Intel® Xeon® CPU E5-2680 v2 @ 2.80GHz 10319.38 67

164 Intel® Xeon® Silver 4116 CPU @ 2.10GHz 10298.34 3

165 AMD EPYC 7281 16-Core Processor 10282.99 2

166 AMD Ryzen Threadripper 1920X 12-Core Processor 10273.37 49

167 Intel® Xeon® CPU E5-2650 v4 @ 2.20GHz 10269.68 9

168 Intel® Xeon® Gold 6134 CPU @ 3.20GHz 10225.37 5

169 AMD Ryzen Threadripper 2950X 16-Core Processor 10201.48 16

170 Intel® Xeon® Platinum 8163 CPU @ 2.50GHz 10174.5 1

171 Intel® Xeon® CPU E5-2663 v3 @ 2.80GHz 10099.79 1

172 Intel® Xeon® Silver 4210R CPU @ 2.40GHz 10098.77 5

173 Intel® Xeon® W-3245 CPU @ 3.20GHz 10077.6 1

174 AMD Ryzen 7 3800XT 8-Core Processor 10066.54 27

175 Intel® Core™ i9-10900K CPU @ 3.70GHz 10024.16 67

176 Genuine Intel® CPU @ 2.00GHz 10012.01 7

177 Intel® Xeon® CPU E5-2687W v3 @ 3.10GHz 9989.91 6

178 Intel® Xeon® CPU E7-4809 v3 @ 2.00GHz 9954.01 1

179 Intel® Xeon® CPU E5-2697 v2 @ 2.70GHz 9901.97 24

180 Intel® Core™ i9-7940X CPU @ 3.10GHz 9889.53 5

181 Intel® Core™ i9-9960X CPU @ 3.10GHz 9841.16 4

182 Genuine Intel® CPU $0000%@ 9816.24 2

183 Genuine Intel® CPU @ 2.80GHz 9801.23 6

184 Genuine Intel® CPU @ 2.50GHz 9732.64 5

185 Intel® Xeon® CPU E5-2698R v4 @ 2.20GHz 9722.24 1

186 Intel® Core™ i9-10900KF CPU @ 3.70GHz 9698.29 13

187 Intel® Core™ i9-10940X CPU @ 3.30GHz 9688.89 9

188 Intel® Xeon® Silver 4114 CPU @ 2.20GHz 9673.99 25

189 Intel® Xeon® CPU E5-2670 v3 @ 2.30GHz 9625.08 13

190 Intel® Xeon® CPU E5-2695 v2 @ 2.40GHz 9618.9 22

191 Intel® Xeon® Silver 4215R CPU @ 3.20GHz 9530.89 2

192 Intel® Xeon® CPU E5-2690 v2 @ 3.00GHz 9528.89 19

193 Intel® Xeon® CPU E7-L8867 @ 2.13GHz 9511.85 1

194 AMD EPYC 7301 16-Core Processor 9509.5 1

195 Intel® Xeon® CPU E5-2667 v4 @ 3.20GHz 9501.1 3

196 Intel® Xeon® Silver 4210 CPU @ 2.20GHz 9483.08 20

197 Intel® Core™ i9-10920X CPU @ 3.50GHz 9310.47 7

198 Intel® Xeon® CPU E5-2640 v4 @ 2.40GHz 9267.67 2

199 Intel® Xeon® CPU E5-2673 v3 @ 2.40GHz 9240.69 16

200 Intel® Xeon® CPU E5-2667 v2 @ 3.30GHz 9218.46 10

201 AMD Ryzen Threadripper 2920X 12-Core Processor 9175.04 3

202 Intel® Xeon® CPU E7- 8870 @ 2.40GHz 9113.11 1

203 Intel® Xeon® Silver 4215 CPU @ 2.50GHz 9099.93 2

204 Intel® Xeon® CPU E5-4620 0 @ 2.20GHz 9063.14 4

205 Intel® Xeon® CPU E5-2687W v2 @ 3.40GHz 9038.32 6

206 Intel® Xeon® CPU E5-2630 v4 @ 2.20GHz 9026.74 32

207 Intel® Xeon® CPU E5-2650 v3 @ 2.30GHz 8956.8 16

208 Intel® Xeon® Silver 4214R CPU @ 2.40GHz 8949.59 4

209 Intel® Xeon® CPU E7- 4850 @ 2.00GHz 8937.43 4

210 Intel® Xeon® CPU E5-2650L v3 @ 1.80GHz 8933.76 10

211 AMD Ryzen 5 5600X 6-Core Processor 8921.88 247

212 Intel® Core™ i9-10850K CPU @ 3.60GHz 8859.75 72

213 Intel® Core™ i9-9940X CPU @ 3.30GHz 8850.81 3

214 AMD Ryzen 5 3600 6-Core Processor 8842.83 869

215 Intel® Xeon® CPU E5-2648L v4 @ 1.80GHz 8730.8 2

216 Intel® Xeon® CPU E5-2660 v4 @ 2.00GHz 8673.78 1

217 Intel® Xeon® Gold 6149 CPU @ 3.10GHz 8653.74 1

218 Intel® Xeon® W-2175 CPU @ 2.50GHz 8169.47 4

219 Intel® Core™ i9-9900K CPU @ 3.60GHz 8145.58 173

220 AMD Ryzen 5 3600XT 6-Core Processor 8145.31 48

221 AMD Ryzen 5 3600X 6-Core Processor 8126.25 154

222 Intel® Xeon® CPU E5-2640 v3 @ 2.60GHz 8096.77 20

223 Intel® Xeon® CPU E5-2670 v2 @ 2.50GHz 8004.61 21

224 Intel® Xeon® Silver 4208 CPU @ 2.10GHz 7910.75 9

225 AMD Opteron™ Processor 6376 7892.91 3

226 Intel® Xeon® CPU E5-2667 v3 @ 3.20GHz 7834.84 6

227 Intel® Xeon® CPU E5-2660 v2 @ 2.20GHz 7814.88 25

228 Intel® Core™ i9-7960X CPU @ 2.80GHz 7758.01 1

229 Intel® Xeon® CPU E5-2470 v2 @ 2.40GHz 7756.03 8

230 Intel® Xeon® CPU E5-2695 v3 @ 2.30GHz 7741.38 2

231 Intel® Xeon® CPU E5-2658 v2 @ 2.40GHz 7712.19 3

232 Intel® Core™ i9-9900X CPU @ 3.50GHz 7678.07 8

233 Intel® Xeon® CPU E5-2698 v4 @ 2.20GHz 7667.89 4

234 AMD Ryzen 9 5900 12-Core Processor 7646.84 1

235 Intel® Xeon® CPU E5-2630 v3 @ 2.40GHz 7644.5 21

236 Intel® Xeon® Silver 4110 CPU @ 2.10GHz 7572.66 19

237 Intel® Core™ i9-7900X CPU @ 3.30GHz 7526.38 23

238 Intel Xeon Processor (Skylake, IBRS) 7504.07 9

239 Intel® Core™ i7-10700K CPU @ 3.80GHz 7492.26 64

240 Intel® Core™ i9-9900KS CPU @ 4.00GHz 7459.37 6

241 Intel® Xeon® Platinum 8280M CPU @ 2.70GHz 7422.86 1

242 Intel® Xeon® CPU E5-2643 v3 @ 3.40GHz 7385.96 5

243 AMD EPYC 7252 8-Core Processor 7375.01 2

244 11th Gen Intel® Core™ i9-11900K @ 3.50GHz 7347.16 9

245 AMD Eng Sample: 100-000000263-30_Y 7300.98 5

246 Intel® Xeon® CPU AWS-1000 @ 2.40GHz 7298.47 1

247 Intel® Xeon® CPU E5-2650 v2 @ 2.60GHz 7277.54 61

248 Intel® Xeon® W-3235 CPU @ 3.30GHz 7247.01 3

249 Intel® Xeon® CPU E5-2687W 0 @ 3.10GHz 7242.23 8

250 Intel® Xeon® CPU E5-2690 0 @ 2.90GHz 7232.8 23

251 Intel® Xeon® CPU AWS-1100 v4 @ 2.40GHz 7224.92 1

252 Intel® Xeon® CPU E5-2643 v2 @ 3.50GHz 7209.33 9

253 11th Gen Intel® Core™ i7-11700K @ 3.60GHz 7160.55 18

254 Intel® Core™ i9-10900X CPU @ 3.70GHz 7159.78 5

255 Intel® Xeon® CPU E5-2699C v4 @ 2.20GHz 7153.23 1

256 Intel® Core™ i9-7920X CPU @ 2.90GHz 7116.27 1

257 AMD EPYC 7232P 8-Core Processor 7020.94 1

258 Intel® Xeon® CPU E5-2620 v4 @ 2.10GHz 7019.6 27

259 AMD Ryzen 7 2700 Eight-Core Processor 6990.95 114

260 Intel® Core™ i7-6950X CPU @ 3.00GHz 6979.1 5

261 Intel® Xeon® CPU E5-2651 v2 @ 1.80GHz 6922.95 8

262 Intel® Xeon® CPU E5-2689 0 @ 2.60GHz 6890.46 40

263 Intel® Xeon® Gold 6254 CPU @ 3.10GHz 6845.61 1

264 Intel® Xeon® CPU E5-2680 0 @ 2.70GHz 6784.81 29

265 Intel® Xeon® CPU E5-2670 0 @ 2.60GHz 6731.64 34

266 AMD Ryzen 7 PRO 2700 Eight-Core Processor 6729.84 5

267 AMD Ryzen 7 2700X Eight-Core Processor 6701.74 204

268 Intel® Xeon® Gold 5217 CPU @ 3.00GHz 6697.83 2

269 Intel® Core™ i9-9900KF CPU @ 3.60GHz 6647.61 26

270 AMD Opteron™ Processor 6220 6593 2