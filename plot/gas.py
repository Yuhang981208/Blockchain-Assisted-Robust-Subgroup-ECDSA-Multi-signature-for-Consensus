import faulthandler;
faulthandler.enable()
import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt

import numpy as np

from pypbc import *

fig, ax = plt.subplots()
node_nums = [3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50]

scheme1 = [706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000,706000]
scheme2 = [406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664,406664]
scheme3 = [576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263,576263]
scheme4 = [201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729,201729]
scheme5 = [111780, 133906, 133906, 156312, 156312, 179064, 179064, 202067, 202067,225379, 225379, 248947, 248947, 272892, 272892, 297081, 297081, 321545,321545, 346234, 346234, 371355, 371355, 396713, 396713, 422363, 422363,448210, 448210, 474523, 474523, 501050, 501050, 527852, 527852, 554827,554827, 582333, 582333, 610029, 610029, 638017, 638017, 666154, 666154,694870, 694870, 723735]


labels = ["BLS-Sakai-CL-RSM", "BLS-RSM", "Sakai-ID-RSM", "BA-RSM", "Single ECDSA"]
plt.rcParams['font.sans-serif'] = ['DejaVu Sans']#指定字体为SimHei

plt.plot(node_nums, scheme1,  label=labels[0])
plt.plot(node_nums, scheme2,   label=labels[1])
plt.plot(node_nums, scheme3,   label=labels[2])
plt.plot(node_nums, scheme4,   label=labels[3])
plt.plot(node_nums, scheme5,   label=labels[4])
ax.ticklabel_format(style='plain')

plt.gcf().subplots_adjust(left=0.15,top=0.9,bottom=0.1)
plt.xlabel("Size of group")  # 横坐标名字
plt.ylabel("Gas consumption")  # 纵坐标名字
plt.legend(fontsize='small')
my_x_ticks = np.arange(0, 51, 5)
my_y_ticks = np.arange(0, 1000001, 100000)
plt.xticks(my_x_ticks)
plt.yticks(my_y_ticks)
fig.savefig('./figures/多重签名验证gas消耗对比.svg', dpi=3200, format='svg')


# nums = [ 5,10,15, 20,25, 30, 35,40,45,50]
# KeyAgg_cost = []
# Combine_cost = []
# for i in range(len(nums)):
#     Combine_cost.append(nums[i] * 19257)
#     KeyAgg_cost.append(nums[i] * 29194)


# fig, ax = plt.subplots()

# plt.rcParams['font.sans-serif'] = ['DejaVu Sans']#指定字体为SimHei

# plt.plot(nums, KeyAgg_cost,  label="KeyAgg cost", markersize = 4, linewidth = 1.0, marker = 's')
# plt.plot(nums, Combine_cost,   label="Combine cost",  markersize = 4, linewidth = 1.0, marker = '*')
# ax.ticklabel_format(style='plain')


# plt.gcf().subplots_adjust(left=0.15,top=0.9,bottom=0.1)
# plt.xlabel("Size of subgroup")  # 横坐标名字
# plt.ylabel("Gas consumption")  # 纵坐标名字
# plt.legend()
# my_x_ticks = np.arange(0, 51, 5)
# my_y_ticks = np.arange(0, 1600001, 200000)
# plt.xticks(my_x_ticks)
# plt.yticks(my_y_ticks)
# fig.savefig('./figures/KeyAgg与Combine合约Gas开销.svg', dpi=3200, format='svg')

# nums = [ 5,10,15, 20,25, 30, 35,40,45,50]
# Paillier_Computation = []
# Paillier_Communication = []
# CL_Computation = []
# CL_Communication = []

# for i in range(len(nums)):
#     Paillier_Computation.append(nums[i] * 400.007 - 399.849)
#     Paillier_Communication.append(nums[i] * 12000 - 11904)
#     CL_Computation.append(nums[i] * 2600.007 - 2599.849)
#     CL_Communication.append(nums[i] * 2000 - 1904)

# fig, ax = plt.subplots()
# print(Paillier_Communication[9], Paillier_Computation[9], CL_Communication[9], CL_Computation[9])

# plt.rcParams['font.sans-serif'] = ['DejaVu Sans']#指定字体为SimHei

# plt.plot(nums, Paillier_Communication,  label="Paillier based scheme", markersize = 4, linewidth = 1.0, marker = 's')
# plt.plot(nums, CL_Communication,   label="CL based scheme",  markersize = 4, linewidth = 1.0, marker = '*')


# plt.gcf().subplots_adjust(left=0.15,top=0.9,bottom=0.1)
# plt.xlabel("Size of subgroup")  # 横坐标名字
# plt.ylabel("communication cost/Bytes")  # 纵坐标名字
# plt.legend()
# my_x_ticks = np.arange(0, 51, 5)
# my_y_ticks = np.arange(0, 600001, 100000)
# plt.xticks(my_x_ticks)
# plt.yticks(my_y_ticks)
# fig.savefig('./figures/通信开销.svg', dpi=3200, format='svg')

# fig, ax = plt.subplots()
# print(Paillier_Communication[9], Paillier_Computation[9], CL_Communication[9], CL_Computation[9])

# plt.rcParams['font.sans-serif'] = ['DejaVu Sans']#指定字体为SimHei

# plt.plot(nums, Paillier_Computation,  label="Paillier based scheme", markersize = 4, linewidth = 1.0, marker = 's')
# plt.plot(nums, CL_Computation,   label="CL based scheme",  markersize = 4, linewidth = 1.0, marker = '*')


# plt.gcf().subplots_adjust(left=0.15,top=0.9,bottom=0.1)
# plt.xlabel("Size of subgroup")  # 横坐标名字
# plt.ylabel("computation cost/ms")  # 纵坐标名字
# plt.legend()
# my_x_ticks = np.arange(0, 51, 5)
# my_y_ticks = np.arange(0, 140001, 20000)
# plt.xticks(my_x_ticks)
# plt.yticks(my_y_ticks)
# fig.savefig('./figures/计算开销.svg', dpi=3200, format='svg')